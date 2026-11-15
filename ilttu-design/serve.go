package main

import (
	"fmt"
	"html"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// answerHeaders are the headers ILTTU's answers' host sends with every file of a restricted answer, the default for a design: the response headers policy's in internal/cdk/answers/edge.go, and the restricted answer's own in internal/edge/origin.go, in asteroid-computing/ilttu.
// An answer gets no Content-Security-Policy, so neither does a preview.
// Strict-Transport-Security is left out, as a preview is plain HTTP on 127.0.0.1.
var answerHeaders = map[string]string{
	"X-Robots-Tag":                 "noindex",
	"X-Content-Type-Options":       "nosniff",
	"Referrer-Policy":              "strict-origin-when-cross-origin",
	"Cross-Origin-Resource-Policy": "same-origin",
	"X-Frame-Options":              "DENY",
	"Cross-Origin-Opener-Policy":   "same-origin",
}

// preview serves a design's folder as ILTTU would serve it as an answer, at the root of its own host, and never lets a browser keep anything.
type preview struct {
	root *os.Root
	log  io.Writer
}

func (p preview) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for name, value := range answerHeaders {
		w.Header().Set(name, value)
	}
	w.Header().Set("Cache-Control", "no-store")
	head := r.Method == http.MethodHead
	if r.Method != http.MethodGet && !head {
		p.notFound(w, r, "ILTTU serves only GET and HEAD")
		return
	}
	name, reason := fileName(r.URL.EscapedPath())
	if reason != "" {
		p.notFound(w, r, reason)
		return
	}
	ct, ok := contentType(name)
	if !ok {
		p.notFound(w, r, name+" is not a file type an answer can hold, so ILTTU refuses it")
		return
	}
	if reason := p.link(name); reason != "" {
		p.notFound(w, r, reason)
		return
	}
	f, err := p.root.Open(name)
	if err != nil {
		p.notFound(w, r, name+" is not in the folder")
		return
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil || !info.Mode().IsRegular() {
		p.notFound(w, r, name+" is not a file")
		return
	}
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(p.log, "%s %s 200\n", r.Method, r.URL.RequestURI())
	if !head {
		io.Copy(w, f)
	}
}

// fileName returns the name, in the folder, of the file the escaped path asks for, with index.html added to a path that ends in /, or why ILTTU would serve nothing at it.
func fileName(escaped string) (string, string) {
	// The edge reads a path with its percent-encodings decoded however many times they were encoded, so each reading is checked.
	reading := escaped
	for range 8 {
		if reason := refused(reading); reason != "" {
			return "", reason
		}
		decoded, err := url.PathUnescape(reading)
		if err != nil {
			return "", "the path is not percent-encoded correctly"
		}
		if decoded == reading {
			break
		}
		reading = decoded
	}
	name, err := url.PathUnescape(escaped)
	if err != nil || !utf8.ValidString(name) {
		return "", "the path is not UTF-8"
	}
	name = strings.TrimPrefix(name, "/")
	if name == "" || strings.HasSuffix(name, "/") {
		name += "index.html"
	}
	if strings.Contains(name, "//") {
		return "", "the path has an empty segment, which no file in an upload has"
	}
	if leftOut(name) {
		return "", name + " is a file a Mac adds, which ILTTU leaves out of an upload"
	}
	return name, ""
}

// refused returns why ILTTU would serve nothing at the path p, read once: a path under /_/, which ILTTU keeps for itself, or one with a . or .. segment or a backslash.
func refused(p string) string {
	if !strings.HasPrefix(p, "/") {
		return "the path does not start with /"
	}
	if strings.Contains(p, `\`) {
		return "the path has a backslash"
	}
	segments := strings.Split(p[1:], "/")
	if segments[0] == "_" {
		return "paths under /_/ are ILTTU's own, and no answer's file is served there"
	}
	for _, s := range segments {
		if s == "." || s == ".." {
			return "the path has a . or .. segment"
		}
	}
	return ""
}

// link returns why ILTTU would refuse name if it, or a folder on its way, is a symbolic link, which ILTTU refuses in an upload.
func (p preview) link(name string) string {
	parts := strings.Split(name, "/")
	for i := range parts {
		at := strings.Join(parts[:i+1], "/")
		info, err := p.root.Lstat(at)
		if err != nil {
			return ""
		}
		if info.Mode()&fs.ModeSymlink != 0 {
			return at + " is a symbolic link, which ILTTU refuses in an upload"
		}
	}
	return ""
}

// notFound answers a 404 that says why, in the page and in the log, so the agent reading either knows what to fix.
func (p preview) notFound(w http.ResponseWriter, r *http.Request, reason string) {
	fmt.Fprintf(p.log, "%s %s 404: %s\n", r.Method, r.URL.RequestURI(), reason)
	body := "<!doctype html>\n<meta charset=\"utf-8\">\n<title>Not found</title>\n<p>Not found: " + html.EscapeString(reason) + ".</p>\n"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(http.StatusNotFound)
	if r.Method != http.MethodHead {
		io.WriteString(w, body)
	}
}

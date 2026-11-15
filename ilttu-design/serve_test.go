package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// design makes a design's folder with the files the cases below ask for.
func design(t *testing.T) *os.Root {
	t.Helper()
	dir := t.TempDir()
	files := map[string]string{
		"index.html":       "<!doctype html><title>Top</title>",
		"styles.css":       "body{}",
		"Page.HTM":         "<p>upper",
		"sub/index.html":   "<p>sub",
		"LICENSE":          "no extension",
		".DS_Store":        "mac",
		"_/index.html":     "ilttu's",
		"__MACOSX/a.html":  "fork",
		"assets/._a.css":   "fork",
		"assets/real.woff": "font",
	}
	for name, body := range files {
		p := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Symlink("styles.css", filepath.Join(dir, "link.css")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("sub", filepath.Join(dir, "linked")); err != nil {
		t.Fatal(err)
	}
	root, err := os.OpenRoot(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { root.Close() })
	return root
}

// TestServe checks each request against what ILTTU's answers' host would answer for the same folder published as an answer.
func TestServe(t *testing.T) {
	h := preview{root: design(t), log: io.Discard}
	cases := []struct {
		method, target string
		status         int
		contentType    string
		body           string
	}{
		{"GET", "/", 200, "text/html; charset=utf-8", "<!doctype html><title>Top</title>"},
		{"HEAD", "/", 200, "text/html; charset=utf-8", ""},
		{"GET", "/index.html", 200, "text/html; charset=utf-8", "<!doctype html><title>Top</title>"},
		{"GET", "/styles.css?v=2", 200, "text/css; charset=utf-8", "body{}"},
		{"GET", "/Page.HTM", 200, "text/html; charset=utf-8", "<p>upper"},
		{"GET", "/sub/", 200, "text/html; charset=utf-8", "<p>sub"},
		{"GET", "/assets/real.woff", 200, "font/woff", "font"},
		{"GET", "/sub", 404, "", ""},
		{"GET", "/missing.html", 404, "", ""},
		{"GET", "/LICENSE", 404, "", ""},
		{"GET", "/.DS_Store", 404, "", ""},
		{"GET", "/__MACOSX/a.html", 404, "", ""},
		{"GET", "/assets/._a.css", 404, "", ""},
		{"GET", "/link.css", 404, "", ""},
		{"GET", "/linked/index.html", 404, "", ""},
		{"GET", "/_/", 404, "", ""},
		{"GET", "/_/index.html", 404, "", ""},
		{"GET", "/%5F/index.html", 404, "", ""},
		{"GET", "/%255F/index.html", 404, "", ""},
		{"GET", "/sub/../index.html", 404, "", ""},
		{"GET", "/sub/%2e%2e/index.html", 404, "", ""},
		{"GET", "/sub/%252E%252E/index.html", 404, "", ""},
		{"GET", "/./index.html", 404, "", ""},
		{"GET", `/sub\index.html`, 404, "", ""},
		{"GET", "/sub%5Cindex.html", 404, "", ""},
		{"GET", "/sub//index.html", 404, "", ""},
		{"GET", "/%ff.html", 404, "", ""},
		{"POST", "/", 404, "", ""},
		{"PUT", "/index.html", 404, "", ""},
	}
	for _, c := range cases {
		t.Run(c.method+" "+c.target, func(t *testing.T) {
			r := httptest.NewRequest(c.method, "http://127.0.0.1"+c.target, nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, r)
			res := w.Result()
			if res.StatusCode != c.status {
				t.Fatalf("status %d, want %d", res.StatusCode, c.status)
			}
			want := map[string]string{
				"Cache-Control":                "no-store",
				"X-Robots-Tag":                 "noindex",
				"X-Content-Type-Options":       "nosniff",
				"Referrer-Policy":              "strict-origin-when-cross-origin",
				"Cross-Origin-Resource-Policy": "same-origin",
				"X-Frame-Options":              "DENY",
				"Cross-Origin-Opener-Policy":   "same-origin",
			}
			for name, value := range want {
				if got := res.Header.Get(name); got != value {
					t.Errorf("%s: %q, want %q", name, got, value)
				}
			}
			for _, name := range []string{"Content-Security-Policy", "Content-Security-Policy-Report-Only", "Strict-Transport-Security", "ETag", "Last-Modified"} {
				if got := res.Header.Get(name); got != "" {
					t.Errorf("%s: %q, want none", name, got)
				}
			}
			if c.status != 200 {
				return
			}
			if got := res.Header.Get("Content-Type"); got != c.contentType {
				t.Errorf("Content-Type: %q, want %q", got, c.contentType)
			}
			body, _ := io.ReadAll(res.Body)
			if string(body) != c.body {
				t.Errorf("body %q, want %q", body, c.body)
			}
		})
	}
}

// TestContentTypes pins a few rows of the table copied from ILTTU's internal/filetypes.
func TestContentTypes(t *testing.T) {
	for name, want := range map[string]string{"a.JS": "text/javascript; charset=utf-8", "a.mjs": "text/javascript; charset=utf-8", "a.svg": "image/svg+xml", "a.woff2": "font/woff2", "a.json": "application/json", "a.wasm": "application/wasm"} {
		if got, ok := contentType(name); !ok || got != want {
			t.Errorf("%s: %q, %v, want %q", name, got, ok, want)
		}
	}
	for _, name := range []string{"Makefile", "a.php", "a.sh", "a.zip", "a."} {
		if got, ok := contentType(name); ok {
			t.Errorf("%s: %q, want none", name, got)
		}
	}
	if len(contentTypes) != 36 {
		t.Errorf("%d types, want ILTTU's 36: copy the table again from internal/filetypes", len(contentTypes))
	}
}

func TestServeNeedsAFolder(t *testing.T) {
	file := filepath.Join(t.TempDir(), "index.html")
	os.WriteFile(file, nil, 0o644)
	for _, args := range [][]string{{"serve"}, {"serve", file}, {"serve", "a", "b"}} {
		if err := run(args, io.Discard, io.Discard); err == nil {
			t.Errorf("%v: no error", args)
		}
	}
	t.Setenv("PORT", "eighty")
	if err := run([]string{"serve", t.TempDir()}, io.Discard, io.Discard); err == nil {
		t.Error("a PORT that is not a number: no error")
	}
}

var _ http.Handler = preview{}

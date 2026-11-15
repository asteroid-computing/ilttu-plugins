package main

import (
	"path"
	"strings"
)

// contentTypes are the files an answer can hold, by extension, and the Content-Type each is served with.
// It is a copy of ILTTU's own table, internal/filetypes in asteroid-computing/ilttu, because answers are served with nosniff and a file served with another type than ILTTU's would act differently once published.
var contentTypes = map[string]string{
	".html":        "text/html; charset=utf-8",
	".htm":         "text/html; charset=utf-8",
	".css":         "text/css; charset=utf-8",
	".js":          "text/javascript; charset=utf-8",
	".mjs":         "text/javascript; charset=utf-8",
	".json":        "application/json",
	".map":         "application/json",
	".webmanifest": "application/manifest+json",
	".png":         "image/png",
	".jpg":         "image/jpeg",
	".jpeg":        "image/jpeg",
	".gif":         "image/gif",
	".webp":        "image/webp",
	".avif":        "image/avif",
	".svg":         "image/svg+xml",
	".ico":         "image/x-icon",
	".woff":        "font/woff",
	".woff2":       "font/woff2",
	".ttf":         "font/ttf",
	".otf":         "font/otf",
	".txt":         "text/plain; charset=utf-8",
	".md":          "text/markdown; charset=utf-8",
	".csv":         "text/csv; charset=utf-8",
	".vtt":         "text/vtt; charset=utf-8",
	".mp3":         "audio/mpeg",
	".wav":         "audio/wav",
	".ogg":         "audio/ogg",
	".oga":         "audio/ogg",
	".m4a":         "audio/mp4",
	".aac":         "audio/aac",
	".flac":        "audio/flac",
	".mp4":         "video/mp4",
	".webm":        "video/webm",
	".ogv":         "video/ogg",
	".mov":         "video/quicktime",
	".wasm":        "application/wasm",
}

// contentType returns the Content-Type of the file name, from its extension in any case, and whether an answer can hold it.
func contentType(name string) (string, bool) {
	t, ok := contentTypes[strings.ToLower(path.Ext(name))]
	return t, ok
}

// leftOut reports whether name is a file a Mac adds, which ILTTU leaves out of an upload: anything under __MACOSX/, a ._ file, and .DS_Store.
func leftOut(name string) bool {
	first, _, _ := strings.Cut(name, "/")
	base := path.Base(name)
	return first == "__MACOSX" || strings.HasPrefix(base, "._") || base == ".DS_Store"
}

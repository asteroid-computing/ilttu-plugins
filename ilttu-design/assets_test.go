package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var (
	defined = regexp.MustCompile(`(?m)^\s*--([a-z0-9-]+):`)
	used    = regexp.MustCompile(`var\(--([a-z0-9-]+)\)`)
	fontSrc = regexp.MustCompile(`url\("([^"]+)"\)`)
)

// TestTokensDefineEveryVariable checks that every var() in components.css and in tokens.css itself names a token tokens.css defines, so the two from the design system stay one set.
func TestTokensDefineEveryVariable(t *testing.T) {
	files, err := templateAssets()
	if err != nil {
		t.Fatal(err)
	}
	have := map[string]bool{}
	for _, m := range defined.FindAllStringSubmatch(string(files["tokens.css"]), -1) {
		have[m[1]] = true
	}
	for _, name := range []string{"components.css", "tokens.css"} {
		for _, m := range used.FindAllStringSubmatch(string(files[name]), -1) {
			if !have[m[1]] {
				t.Errorf("%s uses --%s, which tokens.css does not define", name, m[1])
			}
		}
	}
	for _, m := range fontSrc.FindAllStringSubmatch(string(files["tokens.css"]), -1) {
		if files[m[1]] == nil {
			t.Errorf("tokens.css loads %s, which init does not write", m[1])
		}
	}
}

// TestTokensCSS pins how tokens.json becomes CSS: an alias is a var(), and a themed token's dark value is under prefers-color-scheme: dark.
func TestTokensCSS(t *testing.T) {
	files, err := templateAssets()
	if err != nil {
		t.Fatal(err)
	}
	css := string(files["tokens.css"])
	light, dark, ok := strings.Cut(css, "@media (prefers-color-scheme: dark)")
	if !ok {
		t.Fatal("no dark block")
	}
	for _, want := range []string{"  --ink: oklch(0.23 0.02 165);\n", "  --bg: var(--air);\n", "  --mint-400: var(--mint);\n", "  --space-4: 16px;\n", "  --font-display: \"Gabarito\", ui-sans-serif, system-ui, sans-serif;\n", "  font-weight: 400 900;\n"} {
		if !strings.Contains(light, want) {
			t.Errorf("no %q before the dark block", want)
		}
	}
	for _, want := range []string{"    --bg: var(--pine);\n", "    --action-primary: var(--lemon);\n"} {
		if !strings.Contains(dark, want) {
			t.Errorf("no %q in the dark block", want)
		}
	}
	if strings.Contains(dark, "--ink:") {
		t.Error("an unthemed token in the dark block")
	}
	if regexp.MustCompile(`\{[a-z0-9-]+\}`).MatchString(css) {
		t.Error("an alias left unresolved")
	}
}

// TestInitWritesTheDesignSystem checks that init writes the vendored files byte for byte, and tokens.css as tokensCSS makes it, into the folder's ilttu/, replacing its own files and leaving the rest.
func TestInitWritesTheDesignSystem(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "design")
	if err := os.MkdirAll(filepath.Join(dir, "ilttu"), 0o755); err != nil {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(dir, "index.html"), []byte("mine"), 0o644)
	os.WriteFile(filepath.Join(dir, "ilttu", "tokens.css"), []byte("stale"), 0o644)
	os.WriteFile(filepath.Join(dir, "ilttu", "extra.css"), []byte("mine too"), 0o644)
	for range 2 {
		if err := run([]string{"init", dir}, io.Discard, io.Discard); err != nil {
			t.Fatal(err)
		}
	}
	want := map[string]string{
		"components.css":                               "designsystem/bundle.css",
		"fonts/gabarito-latin-wght-normal.woff2":       "designsystem/fonts/gabarito-latin-wght-normal.woff2",
		"fonts/figtree-latin-wght-normal.woff2":        "designsystem/fonts/figtree-latin-wght-normal.woff2",
		"fonts/jetbrains-mono-latin-wght-normal.woff2": "designsystem/fonts/jetbrains-mono-latin-wght-normal.woff2",
		"fonts/OFL-gabarito.txt":                       "designsystem/fonts/OFL-gabarito.txt",
		"fonts/OFL-figtree.txt":                        "designsystem/fonts/OFL-figtree.txt",
		"fonts/OFL-jetbrains-mono.txt":                 "designsystem/fonts/OFL-jetbrains-mono.txt",
	}
	for name, source := range want {
		got, err := os.ReadFile(filepath.Join(dir, "ilttu", filepath.FromSlash(name)))
		if err != nil {
			t.Fatal(err)
		}
		vendored, err := os.ReadFile(filepath.FromSlash(source))
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, vendored) {
			t.Errorf("%s differs from %s", name, source)
		}
	}
	raw, _ := os.ReadFile("designsystem/tokens.json")
	tokens, err := tokensCSS(raw)
	if err != nil {
		t.Fatal(err)
	}
	if got, _ := os.ReadFile(filepath.Join(dir, "ilttu", "tokens.css")); !bytes.Equal(got, tokens) {
		t.Error("tokens.css is not what tokensCSS makes of tokens.json")
	}
	for name, body := range map[string]string{"index.html": "mine", "ilttu/extra.css": "mine too"} {
		if got, _ := os.ReadFile(filepath.Join(dir, filepath.FromSlash(name))); string(got) != body {
			t.Errorf("%s changed to %q", name, got)
		}
	}
	entries, _ := os.ReadDir(filepath.Join(dir, "ilttu", "fonts"))
	if len(entries) != 6 {
		t.Errorf("%d files in ilttu/fonts, want 6", len(entries))
	}
}

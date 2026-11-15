package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

// designSystem is the vendored copy of the design system, the template assets' one source; designsystem/VENDORED.md says where it came from.
//
//go:embed designsystem/tokens.json designsystem/bundle.css designsystem/fonts
var designSystem embed.FS

// assetsDir is the folder, inside a design's folder, that init writes the template assets to.
const assetsDir = "ilttu"

// templateAssets returns the template assets init writes, by their names under [assetsDir]: tokens.css made from tokens.json, components.css, and the fonts with their licences.
func templateAssets() (map[string][]byte, error) {
	raw, err := designSystem.ReadFile("designsystem/tokens.json")
	if err != nil {
		return nil, err
	}
	tokens, err := tokensCSS(raw)
	if err != nil {
		return nil, fmt.Errorf("reading tokens.json: %w", err)
	}
	components, err := designSystem.ReadFile("designsystem/bundle.css")
	if err != nil {
		return nil, err
	}
	files := map[string][]byte{"tokens.css": tokens, "components.css": components}
	fonts, err := fs.ReadDir(designSystem, "designsystem/fonts")
	if err != nil {
		return nil, err
	}
	for _, f := range fonts {
		b, err := designSystem.ReadFile("designsystem/fonts/" + f.Name())
		if err != nil {
			return nil, err
		}
		files["fonts/"+f.Name()] = b
	}
	return files, nil
}

// writeAssets writes the template assets into dir's [assetsDir] folder, making it if it is not there, and replacing only the files it writes; it returns their paths under dir.
func writeAssets(dir string) ([]string, error) {
	files, err := templateAssets()
	if err != nil {
		return nil, err
	}
	var written []string
	for _, name := range slices.Sorted(maps.Keys(files)) {
		rel := assetsDir + "/" + name
		target := filepath.Join(dir, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return written, err
		}
		if err := os.WriteFile(target, files[name], 0o644); err != nil {
			return written, err
		}
		written = append(written, rel)
	}
	return written, nil
}

// tokenFile is what tokensCSS reads of the design system's tokens.json.
type tokenFile struct {
	Color struct {
		Tokens []token `json:"tokens"`
	} `json:"color"`
	Type struct {
		Fonts []struct {
			Family string `json:"family"`
			File   string `json:"file"`
			Weight string `json:"weight"`
			Style  string `json:"style"`
		} `json:"fonts"`
		Families map[string]string `json:"families"`
	} `json:"type"`
	Spacing tokenGroup `json:"spacing"`
	Radius  tokenGroup `json:"radius"`
	Shadow  tokenGroup `json:"shadow"`
	Size    tokenGroup `json:"size"`
	Layout  tokenGroup `json:"layout"`
}

type tokenGroup struct {
	Tokens []token `json:"tokens"`
}

// token is one design token: its value is a CSS value, a {name} alias of another token, or, for a token that differs by theme, an object of the light and the dark value.
type token struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

// alias is a token value that names another token, such as {ink}.
var alias = regexp.MustCompile(`^\{([a-z0-9-]+)\}$`)

// tokensCSS makes tokens.css from the design system's tokens.json: an @font-face for each font, every token as a custom property on :root, and the dark value of each themed token under prefers-color-scheme: dark.
func tokensCSS(raw []byte) ([]byte, error) {
	var t tokenFile
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, err
	}
	var b strings.Builder
	b.WriteString("/* The ilttu design system's tokens, written by ilttu-design init from its tokens.json. Run init again rather than editing this file. */\n\n")
	for _, f := range t.Type.Fonts {
		fmt.Fprintf(&b, "@font-face {\n  font-family: %q;\n  font-style: %s;\n  font-display: swap;\n  font-weight: %s;\n  src: url(%q) format(\"woff2\");\n}\n", f.Family, f.Style, f.Weight, f.File)
	}
	var dark []string
	b.WriteString("\n:root {\n")
	for _, name := range slices.Sorted(maps.Keys(t.Type.Families)) {
		fmt.Fprintf(&b, "  --font-%s: %s;\n", name, t.Type.Families[name])
	}
	groups := [][]token{t.Color.Tokens, t.Spacing.Tokens, t.Radius.Tokens, t.Shadow.Tokens, t.Size.Tokens, t.Layout.Tokens}
	for _, group := range groups {
		for _, tok := range group {
			var plain string
			var themed struct{ Light, Dark string }
			switch {
			case json.Unmarshal(tok.Value, &plain) == nil:
				fmt.Fprintf(&b, "  --%s: %s;\n", tok.Name, cssValue(plain))
			case json.Unmarshal(tok.Value, &themed) == nil && themed.Light != "" && themed.Dark != "":
				fmt.Fprintf(&b, "  --%s: %s;\n", tok.Name, cssValue(themed.Light))
				dark = append(dark, fmt.Sprintf("    --%s: %s;\n", tok.Name, cssValue(themed.Dark)))
			default:
				return nil, fmt.Errorf("token %q has a value that is neither a string nor a light and a dark one", tok.Name)
			}
		}
	}
	b.WriteString("}\n")
	if len(dark) > 0 {
		b.WriteString("\n@media (prefers-color-scheme: dark) {\n  :root {\n")
		for _, line := range dark {
			b.WriteString(line)
		}
		b.WriteString("  }\n}\n")
	}
	return []byte(b.String()), nil
}

// cssValue returns a token's value as CSS: an alias becomes the var() of the token it names.
func cssValue(v string) string {
	if m := alias.FindStringSubmatch(v); m != nil {
		return "var(--" + m[1] + ")"
	}
	return v
}

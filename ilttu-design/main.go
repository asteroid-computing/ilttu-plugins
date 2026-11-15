// Command ilttu-design previews a design's folder on localhost as ILTTU will serve it as an answer, and copies the ilttu design system's template assets into it.
//
//	ilttu-design serve [-port N] <dir>   serve <dir> on 127.0.0.1 at -port, else $PORT, else a free port
//	ilttu-design init <dir>              write the template assets into <dir>/ilttu/
//	ilttu-design version                 print the version
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// version is the release's tag, set when a release is built.
var version = "dev"

const usage = `ilttu-design previews a design on localhost as ILTTU will serve it as an answer.

  ilttu-design serve [-port N] <dir>   serve <dir> on 127.0.0.1 at -port, else $PORT, else a free port
  ilttu-design init <dir>              write the ilttu design system's assets into <dir>/ilttu/
  ilttu-design version                 print the version
`

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "ilttu-design:", err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return errors.New("name a command")
	}
	switch args[0] {
	case "serve":
		return serve(args[1:], stdout, stderr)
	case "init":
		return initAssets(args[1:], stdout, stderr)
	case "version":
		fmt.Fprintln(stdout, version)
		return nil
	case "help", "-h", "--help":
		fmt.Fprint(stdout, usage)
		return nil
	}
	fmt.Fprint(stderr, usage)
	return fmt.Errorf("no command %q", args[0])
}

func serve(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("serve", flag.ContinueOnError)
	flags.SetOutput(stderr)
	port := flags.Int("port", -1, "the port to listen on; $PORT, else a free one, when left out")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 1 {
		return errors.New("serve takes one folder: ilttu-design serve [-port N] <dir>")
	}
	dir, err := folder(flags.Arg(0))
	if err != nil {
		return err
	}
	if *port < 0 {
		*port = 0
		if p := os.Getenv("PORT"); p != "" {
			if *port, err = strconv.Atoi(p); err != nil {
				return fmt.Errorf("PORT %q is not a number", p)
			}
		}
	}
	root, err := os.OpenRoot(dir)
	if err != nil {
		return err
	}
	defer root.Close()
	if _, err := root.Stat("index.html"); err != nil {
		fmt.Fprintf(stderr, "ilttu-design: %s has no index.html at its top, which an answer needs\n", dir)
	}
	l, err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(*port)))
	if err != nil {
		return err
	}
	fmt.Fprintf(stdout, "Serving %s at http://127.0.0.1:%d/\n", dir, l.Addr().(*net.TCPAddr).Port)
	return http.Serve(l, preview{root: root, log: stderr})
}

func initAssets(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("init", flag.ContinueOnError)
	flags.SetOutput(stderr)
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 1 {
		return errors.New("init takes one folder: ilttu-design init <dir>")
	}
	dir, err := filepath.Abs(flags.Arg(0))
	if err != nil {
		return err
	}
	written, err := writeAssets(dir)
	if err != nil {
		return err
	}
	for _, name := range written {
		fmt.Fprintln(stdout, "wrote", filepath.Join(dir, filepath.FromSlash(name)))
	}
	fmt.Fprintf(stdout, "\nLink them in the <head> of each page at the folder's top:\n\n  <link rel=\"stylesheet\" href=\"%[1]s/tokens.css\">\n  <link rel=\"stylesheet\" href=\"%[1]s/components.css\">\n", assetsDir)
	return nil
}

// folder returns the absolute path of dir, which must be a folder.
func folder(dir string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("%s is not a folder", abs)
	}
	return abs, nil
}

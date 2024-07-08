package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rcrowley/mergician/html"
	"golang.org/x/net/html/atom"
)

func init() {
	log.SetFlags(0)
}

func main() {
	document := flag.Bool("d", false, "wrap the fragment in a complete HTML document")
	output := flag.String("o", "-", "write to this file instead of standard output")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: frag [-d] [-o <output>] <tag> <input>
  -d           wrap the fragment in a complete HTML document")
  -o <output>  write to this file instead of standard output
  <tag>        tag (optionally with attributes) at the root of the fragment to extract
  <input>      pathname to an input HTML file
`)
	}
	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	in := must2(html.ParseFile(flag.Arg(1)))

	tag := must2(html.ParseString(flag.Arg(0)))

	out := html.Find(in, html.Match(tag))
	if out == nil {
		log.Fatalf("%s not found", flag.Arg(0))
	}
	if *document {
		out = wrap(out)
	}

	var w io.Writer
	if *output == "-" {
		w = os.Stdout
	} else {
		f := must2(os.Create(*output))
		defer f.Close()
		w = f
	}
	must(html.Render(w, out))

	// Fragments rooted in an element (as opposed to bare text nodes) can't
	// end with a trailing newline so we add one because we are a well-
	// behaved Unix program. This isn't included in html.Render because that
	// function isn't necessarily always used as the output of a Unix program.
	if out.Type == html.ElementNode {
		fmt.Fprintln(w, "")
	}

}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func must2[T any](v T, err error) T {
	must(err)
	return v
}

func wrap(in *html.Node) *html.Node {
	out := must2(html.ParseString(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
</head>
<body>
</body>
</html>
`))
	html.Find(out, html.IsAtom(atom.Body)).AppendChild(html.CopyNode(in))
	return out
}

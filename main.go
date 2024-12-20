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

func Main(args []string, stdin io.Reader, stdout io.Writer) {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	document := flags.Bool("d", false, "wrap the fragment in a complete HTML document")
	inner := flags.Bool("i", false, "unwrap the fragment to leave only its inner HTML")
	output := flags.String("o", "-", "write to this file instead of standard output")
	flags.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: frag [-d|-i] [-o <output>] <tag> <input>
  -d           wrap the fragment in a complete HTML document
  -i           unwrap the fragment to leave only its inner HTML
  -o <output>  write to this file instead of standard output
  <tag>        tag (optionally with attributes) at the root of the fragment to extract
  <input>      input HTML file

Synopsis: frag extracts document fragments from complete HTML documents.
`)
	}
	flags.Parse(args[1:])
	if *document && *inner || flags.NArg() != 2 {
		flags.Usage()
		os.Exit(1)
	}

	var w io.Writer
	if *output == "-" {
		w = stdout
	} else {
		f := must2(os.Create(*output))
		defer f.Close()
		w = f
	}

	in := must2(html.ParseFile(flags.Arg(1)))

	tag := must2(html.ParseString(flags.Arg(0)))

	out := html.Find(in, html.Match(tag))
	if out == nil {
		log.Fatalf("%s not found", flags.Arg(0))
	}
	if *document {
		must(html.Render(w, wrap(out)))
	} else if *inner {
		for n := out.FirstChild; n != nil; n = n.NextSibling {
			if (n != out.FirstChild && n != out.LastChild) || !html.IsWhitespace(n) { // strip leading/trailing whitespace
				must(html.Render(w, n))
			}
		}
	} else {
		must(html.Render(w, out))
	}

	// Fragments rooted in an element (as opposed to bare text nodes) can't
	// end with a trailing newline so we add one because we are a well-
	// behaved Unix program. This isn't included in html.Render because that
	// function isn't necessarily always used as the output of a Unix program.
	if out.Type == html.ElementNode {
		fmt.Fprintln(w, "")
	}

}

func init() {
	log.SetFlags(0)
}

func main() {
	Main(os.Args, os.Stdin, os.Stdout)
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

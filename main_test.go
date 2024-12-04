package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainH1(t *testing.T) {
	stdout := &bytes.Buffer{}
	Main([]string{"frag", "<h1>", "test.html"}, os.Stdin, stdout)
	if s := stdout.String(); s != "<h1>Hello, world!</h1>\n" {
		t.Fatal(s)
	}
}

func TestMainTitle(t *testing.T) {
	stdout := &bytes.Buffer{}
	Main([]string{"frag", "<title>", "test.html"}, os.Stdin, stdout)
	if s := stdout.String(); s != "<title>Hello, world!</title>\n" {
		t.Fatal(s)
	}
}

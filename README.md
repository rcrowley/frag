Frag
====

TODO

Installation
------------

```sh
go get github.com/rcrowley/frag
```

Usage
-----

```sh
frag [-d|-i] [-o <output>] <tag> <input>
```

* `-d`: wrap the fragment in a complete HTML document
* `-i`: unwrap the fragment to leave only its inner HTML
* `-o <output>`: write to this file instead of standard output
* `<tag>`: tag (optionally with attributes) at the root of the fragment to extract
* `<input>`: input HTML file

See also
--------

Frag is part of the [Mergician](https://github.com/rcrowley/mergician) suite of tools that manipulate HTML documents:

* [Deadlinks](https://github.com/rcrowley/deadlinks): Scan a document root directory for dead links
* [Electrostatic](https://github.com/rcrowley/electrostatic): Mergician-powered, pure-HTML CMS
* [Sitesearch](https://github.com/rcrowley/sitesearch): Index a document root directory and serve queries to it in AWS Lambda

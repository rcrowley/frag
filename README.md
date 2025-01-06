Frag
====

Extract fragments of HTML documents.

For example, given this `test.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Hello, world!</title>
</head>
<body>
<article>
<h1>Hello, world!</h1>
<p>Lovely day for a test, isn&rsquo;t it?</p>
</article>
</body>
</html>
```

And this use of Frag:

```sh
frag '<h1>' test.html
```

You get this fragment of the HTML document:

```
<h1>Hello, world!</h1>
```

Installation
------------

```sh
go install github.com/rcrowley/frag@latest
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
* [Feed](https://github.com/rcrowley/feed): Scan a document root directory to construct an Atom feed
* [Sitesearch](https://github.com/rcrowley/sitesearch): Index a document root directory and serve queries to it in AWS Lambda

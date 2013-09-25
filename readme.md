# goserve

An extremely simple static web server in [go][go].

## Quick start

1. Open shell
2. Type: `go get github.com/jugglingnutcase/goserve`.
3. Browse to a folder with stuff in it
4. Type: `goserve`

## Installation

1. Install [go][go] if you don't already have it
2. Set up your go path like mentioned in [How to Write Go code](http://golang.org/doc/code.html).
3. Run `go get github.com/jugglingnutcase/goserve`
4. [Optional] You may want to put go binaries on your path for easy access.

## Usage

Assuming that go executables are accessible via your path, in the directory you want to serve, type: `goserve`

Your files will be served at http://localhost:4000/.

If you want to specify a different port, give a port argument to the application: `goserve -port 3000`

## Why?

i did this while learning go. i know this is trivial, but i figured i'd make something i can actually use while i was doing trivial examples. i'm sure there's a million others (or not because it's so darn simple to do), but i've never found them. Also, i couldn't find a way to deploy Go code via Gist, because to me this makes more sense as a Gist.

[go]:http://golang.org


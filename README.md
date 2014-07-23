[![Build Status](https://img.shields.io/travis/bjeanes/go-lifx.svg)](https://travis-ci.org/bjeanes/go-lifx)
[![GoDoc](https://godoc.org/github.com/bjeanes/go-lifx?status.svg)](http://godoc.org/github.com/bjeanes/go-lifx)
[![License](http://img.shields.io/badge/license-MIT-green.svg)](http://bjeanes.mit-license.org/)

go-lifx
=======

This repository contains (or will contain) Go code for programmatically dealing
with LIFX light bulbs on the local network.

It is (or will be) comprised of loosely three sections:

* **Protocol** - Encoding/decoding of LIFX messages and transmission/reception
  of such.
* **Client** - Higher-level abstractions ("light bulb", "label", "colors",
  etc.) and actions ("set label", "set lights to a specific colour", etc.).
* **CLI Tools** - These are foremost examples of using the **Client** and/or
  **Protocol** components but also serve as debugging tools and command line
  control over bulbs (e.g. from Bash scripts).

## Command Line Tools

There are currently two CLI tools, with their own READMEs:

* [`lifx`](/cli/lifx/README.md)
* [`lifx-snoop`](/cli/lifx-snoop/README.md)

## License

The MIT License (MIT)
Copyright © 2014 Bo Jeanes <me@bjeanes.com>

[Full license here](http://bjeanes.mit-license.org/).

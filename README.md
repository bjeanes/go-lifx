go-lifx
=======

This repository contains (or will contain) Go code for programmatically dealing
with LIFX light bulbs on the local network.

It is (or will be) comprised of loosely three sections:

* **Protocol** - Encoding/decoding of LIFX messages and transmission/reception
  of such.
* **Library** - Higher-level abstractions ("light bulb", "label", "colors",
  etc.) and actions ("set label", "set lights to a specific colour", etc.).
* **CLI Tools** - These are foremost examples of using the **Library** and/or
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

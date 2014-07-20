# Type Registry

This is a hacky experiment that I'll probably back out of, but...

The LIFX protocol encodes the type of message using a uint in the header. There
are more than a hundred message types so the pervasive `switch` statements were
getting unwieldy. This is an attempt to be able to map between a struct and its
message ID.

Likely, I'll replace this questionable strategy with some interfaces and generate
the boilerplates from a more concise schema of some kind.

# `lifx`

*This is mostly unimplemented so this README currently represents "readme driven development"*

## Install

``` sh-session
$ go install github.com/bjeanes/go-lifx/cli/lifx
$ PATH="`go env GOPATH`/bin:$PATH" # probably in your .bashrc or equivalent
```

## Usage


``` sh-session
$ lifx -h
NAME:
   lifx - Control LIFX bulbs from the command line

USAGE:
   lifx [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   list                List all discovered bulbs
   on                  Turn light(s) on
   off                 Turn light(s) off
   change              Change specified color attributes of light(s)
   help, h             Show a list of commands or help for one command

GLOBAL OPTIONS:
   --tag, -t '*'        Limit actions to bulbs with specified tag
   --label, -l '*'      Limit actions to bulb with specified label
   --help, -h           Show help
   --version, -v        Print the version

$ lifx change -h
NAME:
   change - Change specified color attributes of light(s)

USAGE:
   command change [command options] [arguments...]

DESCRIPTION:


OPTIONS:
   --hue '0.85'             Hue (0.0-360.0)
   --brightness '0.85'      Brightness (0.0-1.0)
   --saturation '0.85'      Saturation (0.0-1.0)
   --kelvin, --temp '5000'  Temperature in °K (2500-10000)
   --color 'white'          Set to named color
   --colors                 List available named colors
```

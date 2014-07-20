go-lifx
=======

## `lifx`

Basically unimplemented, but the idea will be:

``` sh-session
$ go run lifx.go -h
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

$ go run lifx.go change -h
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

## `lifx-snoop`

``` sh-session
$ go run lifx-snoop.go
LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3700} Dim:0 Power:65535 Label:[240 159 147 136 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] Tags:1}
LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:65535 Kelvin:3700} Dim:0 Power:0 Label:[240 159 145 167 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] Tags:0}
LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:65535 Kelvin:3000} Dim:0 Power:0 Label:[240 159 145 166 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] Tags:0}
LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3700} Dim:0 Power:65535 Label:[240 159 147 154 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] Tags:1}
LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3700} Dim:0 Power:65535 Label:[240 159 142 163 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] Tags:1}
LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3200} Dim:0 Power:65535 Label:[240 159 148 173 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] Tags:1}
LIFXMessage(deviceGetPanGateway){}
LIFXMessage(deviceStatePanGateway){Service:2 Port:0}
LIFXMessage(deviceStatePanGateway){Service:1 Port:56700}
LIFXMessage(deviceStateTagLabels){Tags:1 Label:[72 111 109 101 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]}
```

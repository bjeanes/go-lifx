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
DATA: length=88
      000  58 00 00 54 00 00 00 00  d0 73 d5 00 14 cd 00 00  |X..T.....s......|
      010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
      020  6b 00 00 00 00 00 00 00  ee 6d 74 0e 00 00 ff ff  |k........mt.....|
      030  f0 9f 93 88 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      050  01 00 00 00 00 00 00 00                           |........|
MSG:  LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3700} Dim:0 Power:65535 Label:📈 Tags:1}

DATA: length=88
      000  58 00 00 54 00 00 00 00  d0 73 d5 00 03 c4 00 00  |X..T.....s......|
      010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
      020  6b 00 00 00 00 00 00 00  ff ff 74 0e 00 00 00 00  |k.........t.....|
      030  f0 9f 91 a7 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      050  00 00 00 00 00 00 00 00                           |........|
MSG:  LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:65535 Kelvin:3700} Dim:0 Power:0 Label:👧 Tags:0}

DATA: length=88
      000  58 00 00 54 00 00 00 00  d0 73 d5 00 fb b0 00 00  |X..T.....s......|
      010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
      020  6b 00 00 00 00 00 00 00  ee 6d 74 0e 00 00 ff ff  |k........mt.....|
      030  f0 9f 93 9a 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      050  01 00 00 00 00 00 00 00                           |........|
MSG:  LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3700} Dim:0 Power:65535 Label:📚 Tags:1}

DATA: length=88
      000  58 00 00 54 00 00 00 00  d0 73 d5 00 0d 76 00 00  |X..T.....s...v..|
      010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
      020  6b 00 00 00 00 00 00 00  ff ff b8 0b 00 00 00 00  |k...............|
      030  f0 9f 91 a6 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      050  00 00 00 00 00 00 00 00                           |........|
MSG:  LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:65535 Kelvin:3000} Dim:0 Power:0 Label:👦 Tags:0}

DATA: length=88
      000  58 00 00 54 00 00 00 00  d0 73 d5 00 f3 26 00 00  |X..T.....s...&..|
      010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
      020  6b 00 00 00 00 00 00 00  ee 6d 74 0e 00 00 ff ff  |k........mt.....|
      030  f0 9f 8e a3 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      050  01 00 00 00 00 00 00 00                           |........|
MSG:  LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3700} Dim:0 Power:65535 Label:🎣 Tags:1}

DATA: length=88
      000  58 00 00 54 00 00 00 00  d0 73 d5 00 0c af 00 00  |X..T.....s......|
      010  4c 49 46 58 56 32 00 00  00 00 00 00 00 00 00 00  |LIFXV2..........|
      020  6b 00 00 00 00 00 00 00  ee 6d 80 0c 00 00 ff ff  |k........m......|
      030  f0 9f 94 ad 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
      050  01 00 00 00 00 00 00 00                           |........|
MSG:  LIFXMessage(lightState){Color:{Hue:0 Saturation:0 Brightness:28142 Kelvin:3200} Dim:0 Power:65535 Label:🔭 Tags:1}
```

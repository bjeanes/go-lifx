# `lifx-snoop`

This is a tool that can be used to debug LIFX protocol UDP traffic on the network.

It displays the bytes of all received UDP datagrams as well as the attempted decoding of the
datagram into a known LIFX message. Decoding errors are also output.

## Install

``` sh-session
$ go install github.com/bjeanes/go-lifx/cli/lifx-snoop
$ PATH="`go env GOPATH`/bin:$PATH" # probably in your .bashrc or equivalent
```

## Usage

``` sh-session
$ lifx-snoop
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

## TODO

* Filtering by type (light, device, sensor, wifi, etc)
* Filtering by device (specific light, tabel, gateway, etc)
* Suppressing errors
* Hiding the raw data or showing it in different formats
* Logging session transcripts
* Replaying session transcripts to the network
  * At this point `lifx-debug` might be a better name, though!

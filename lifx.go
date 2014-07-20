package main

import (
	"github.com/bjeanes/go-lifx/version"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lifx"
	app.Usage = "Control LIFX bulbs from the command line"
	app.Version = version.Version
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{"tag, t", "*", "Limit action to bulbs with specified tag"},
		cli.StringFlag{"label, l", "*", "Limit action to bulb with specified label"},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:   "list",
			Usage:  "List all discovered bulbs",
			Action: func(_ *cli.Context) {},
		},
		cli.Command{
			Name:   "on",
			Usage:  "Turn light(s) on",
			Action: func(_ *cli.Context) {},
		},
		cli.Command{
			Name:   "off",
			Usage:  "Turn light(s) off",
			Action: func(_ *cli.Context) {},
		},
		cli.Command{
			Name:  "change",
			Usage: "Change specified color attributes of light(s)",
			Flags: []cli.Flag{
				cli.Float64Flag{"hue", 0.85, "Hue (0.0-360.0)"},
				cli.Float64Flag{"brightness", 0.85, "Brightness (0.0-1.0)"},
				cli.Float64Flag{"saturation", 0.85, "Saturation (0.0-1.0)"},
				cli.IntFlag{"kelvin, temp", 5000, "Temperature in °K (2500-10000)"},
				cli.StringFlag{"color", "white", "Set to named color"},
				cli.BoolFlag{"colors", "List available named colors"},
			},
			Action: func(_ *cli.Context) {},
		},
	}

	app.Run(os.Args)
}

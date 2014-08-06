package main

import (
	"fmt"
	"github.com/bjeanes/go-lifx"
	"github.com/bjeanes/go-lifx/client"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lifx"
	app.Usage = "Control LIFX bulbs from the command line"
	app.Version = lifx.Version
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "tag, t", Value: "*", Usage: "Limit action to bulbs with specified tag"},
		cli.StringFlag{Name: "label, l", Value: "*", Usage: "Limit action to bulb with specified label"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "List all discovered bulbs",
			Action: func(_ *cli.Context) {
				client := client.New()
				fmt.Println("Discovered bulbs:")
				for light := range client.Discover() {
					fmt.Println("  " + light.Label())
				}
			},
		},
		{
			Name:   "on",
			Usage:  "Turn light(s) on",
			Action: func(_ *cli.Context) {},
		},
		{
			Name:   "off",
			Usage:  "Turn light(s) off",
			Action: func(_ *cli.Context) {},
		},
		{
			Name:  "change",
			Usage: "Change specified color attributes of light(s)",
			Flags: []cli.Flag{
				cli.Float64Flag{Name: "hue, H", Value: 0.85, Usage: "Hue (0.0-360.0)"},
				cli.Float64Flag{Name: "brightness, B", Value: 0.85, Usage: "Brightness (0.0-1.0)"},
				cli.Float64Flag{Name: "saturation, S", Value: 0.85, Usage: "Saturation (0.0-1.0)"},
				cli.IntFlag{Name: "kelvin, temp, K", Value: 5000, Usage: "Temperature in °K (2500-10000)"},
				cli.StringFlag{Name: "color", Value: "white", Usage: "Set to named color"},
				cli.BoolFlag{Name: "colors", Usage: "List available named colors"},
			},
			Action: func(_ *cli.Context) {},
		},
	}

	app.Run(os.Args)
}

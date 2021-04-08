package main

import (
	"log"
	"os"

	"github.com/loganprice/vaccine-alert/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Covid Alert",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "zip-code",
				Usage:    "Zip Code to be searched",
				EnvVars:  []string{"ZIPCODE", "ZIP"},
				Required: true,
			},
			&cli.IntFlag{
				Name:    "max-distance",
				Value:   25,
				Usage:   "Distance from Zip Code to search",
				EnvVars: []string{"DISTANCE"},
			},
			&cli.StringFlag{
				Name:     "zip-key",
				Usage:    "API Key for https://www.zipcodeapi.com/",
				EnvVars:  []string{"ZIP_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "ifttt-key",
				Usage:    "API Key for https://ifttt.com/",
				EnvVars:  []string{"IFTTT_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "ifttt-event",
				Usage:    "Event ID for https://ifttt.com/ webhook",
				EnvVars:  []string{"IFTTT_KEY"},
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "jj-allowed",
				Usage:   "Show alerts for Johnson and Johnson Vaccinne",
				EnvVars: []string{"JJ"},
				Value:   true,
			},
			&cli.BoolFlag{
				Name:    "pfizer-allowed",
				Usage:   "Show alerts for Pfizer Vaccinne",
				EnvVars: []string{"PFIZER"},
				Value:   true,
			},
			&cli.BoolFlag{
				Name:    "moderna-allowed",
				Usage:   "Show alerts for Moderna Vaccinne",
				EnvVars: []string{"MODERNA"},
				Value:   true,
			},
			&cli.IntFlag{
				Name:    "interval",
				Usage:   "Interval in seconds to to check for appointments",
				EnvVars: []string{"INTERVAL"},
				Value:   15,
			},
		},
		Action: cmd.Run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

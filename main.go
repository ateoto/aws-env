package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/vaughan0/go-ini"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "aws-env"
	app.Version = "0.0.1"
	app.Usage = "Set AWS EnvVars from .aws/config file"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "profile, p",
			Value: "default",
			Usage: "Specify a profile to use",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: filepath.Join(os.Getenv("HOME"), ".aws", "config"),
			Usage: "Specify a config file to use",
		},
	}

	app.Action = func(c *cli.Context) {
		profileKey := c.String("profile")

		if c.String("profile") != "default" {
			profileKey = fmt.Sprintf("profile %s", c.String("profile"))
		}

		file, err := ini.LoadFile(c.String("config"))
		if err != nil {
			log.Fatal(err)
		}

		for key, value := range file[profileKey] {
			err := os.Setenv(strings.ToUpper(key), value)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("export %s=%s\n", strings.ToUpper(key), value)
		}
	}

	app.Run(os.Args)
}

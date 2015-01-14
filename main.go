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
	app.Version = "1.0.0"
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
		cli.BoolFlag{
			Name:  "access-only, a",
			Usage: "Will only print the access key, useful for scripting.",
		},
		cli.BoolFlag{
			Name:  "secret-only, s",
			Usage: "Will only print the secret key, useful for scripting.",
		},
	}

	app.Action = func(c *cli.Context) {
		profileKey := c.String("profile")
		access_only := c.String("access-only")
		secret_only := c.String("secret-only")

		if c.String("profile") != "default" {
			profileKey = fmt.Sprintf("profile %s", c.String("profile"))
		}

		file, err := ini.LoadFile(c.String("config"))
		if err != nil {
			log.Fatal(err)
		}

		aws_config := make(map[string]string)
		for key, value := range file[profileKey] {
			err := os.Setenv(strings.ToUpper(key), value)
			if err != nil {
				log.Fatal(err)
			}
			aws_config[key] = value
		}

		if access_only == "true" {
			fmt.Printf("%s", aws_config["aws_access_key_id"])
		} else if secret_only == "true" {
			fmt.Printf("%s", aws_config["aws_secret_access_key"])
		} else {
			for key, value := range aws_config {
				fmt.Printf("export %s=%s\n", strings.ToUpper(key), value)
			}
		}
	}

	app.Run(os.Args)
}

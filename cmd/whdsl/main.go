package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/maxihafer/whdsl/cmd/whdsl/internal"
)

func main() {
	app := cli.App{
		Name: "whdsl-service",
		Description: "Manage your inventory like a boss",
		Commands: []*cli.Command{
			{
				Name: "server",
				Usage: "Start the grpc server and listen for calls",
				Action: func(c *cli.Context) error {
					s, err := internal.NewServerFromEnv()

					if err != nil {
						return err
					}

					return s.Run()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"

	"github.com/maxihafer/whdsl/cmd/whdsl/internal"
	"github.com/maxihafer/whdsl/pkg/emitter"
)

func main() {
	app := cli.App{
		Name:        "whdsl-service",
		Description: "Manage your inventory like a boss",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start the grpc server and listen for calls",
				Action: func(c *cli.Context) error {
					s, err := internal.NewServerFromEnv()

					if err != nil {
						return err
					}

					return s.Run()
				},
			},
			{
				Name:  "emitter",
				Usage: "Start an emitter, providing demo data for the service",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "config",
						Aliases:  []string{"c"},
						Required: true,
					},
				},
				Action: func(context *cli.Context) error {
					path := context.Path("config")

					bytes, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					emitterConfig := &emitter.Config{}
					err = yaml.Unmarshal(bytes,emitterConfig)
					if err != nil {
						return err
					}

					c, err := internal.NewClientFromEnv(emitterConfig)
					if err != nil {
						return err
					}

					return c.Run(context.Context)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

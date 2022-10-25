package main

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"whdsl/cmd/whdsl/internal/emitter"
	"whdsl/cmd/whdsl/internal/server"
)

func main() {
	ctx := context.Background()

	app := cli.App{
		Name: "whdsl",
		Description: "inventory management software for whdsl",
		Commands: []*cli.Command{
			{
				Name: "serve",
				Usage: "serve the backend api for inventory management",
				Action: func(c *cli.Context) error {
					svc := server.Server{}

					return svc.Run(ctx)
				},
			},
			{
				Name: "emitter",
				Usage: "start the random transaction emitter",
				Action: func(c *cli.Context) error {
					svc := emitter.Server{}

					defer func() {
						err := svc.Close(ctx)
						if err != nil {
							logrus.Fatal(err)
						}
					}()

					return svc.Run(ctx)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

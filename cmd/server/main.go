package main

import (
	"github.com/alexpfx/go_evrelay/internal/evrelay"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   50051,
			},
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"u"},
				Value:   "0.0.0.0",
			},
		},
		Action: func(context *cli.Context) error {
			port := context.Int("port")
			host := context.String("host")

			server := evrelay.NewServer(host, port)
			server.Start()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

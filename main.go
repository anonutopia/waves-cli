package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

var wnc *WavesNodeClient

func main() {
	wnc = &WavesNodeClient{}
	wnc.Host = DEFAULT_HOST
	wnc.Port = DEFAULT_PORT

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api-key",
				Usage: "API key for waves node",
			},
		},
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "activation",
				Usage: "node activation commands",
				Subcommands: []*cli.Command{
					{
						Name:   "status",
						Usage:  "get activation status",
						Action: CommandHandler,
					},
				},
			},
			{
				Name: "node",
				// Aliases: []string{"a"},
				Usage: "node commands",
				Subcommands: []*cli.Command{
					{
						Name:   "version",
						Usage:  "get node version",
						Action: CommandHandler,
					},
					{
						Name:   "stop",
						Usage:  "stop waves node",
						Action: CommandHandler,
					},
					{
						Name:   "status",
						Usage:  "get node status",
						Action: CommandHandler,
					},
				},
			},
			{
				Name:   "addresses",
				Usage:  "list wallet addresses",
				Action: CommandHandler,
				// Subcommands: []*cli.Command{
				// 	{
				// 		Name:   "status",
				// 		Usage:  "get activation status",
				// 		Action: CommandHandler,
				// 	},
				// },
			},
		},
		Before: func(c *cli.Context) error {
			wnc.ApiKey = c.String("api-key")
			return nil
		},
	}

	app.Run(os.Args)
}

package main

import (
	"bufio"
	"fmt"
	"github.com/alexpfx/go_evrelay/internal/evrelay/client"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sync"
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

		Commands: []*cli.Command{
			{
				Name: "stats",
				Action: func(ctx *cli.Context) error {
					host := ctx.String("host")
					port := ctx.Int("port")

					c := client.NewStat(host, port)
					go c.Start()

					readCmd(&c)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func readCmd(cl *client.Stats) {
	var wg sync.WaitGroup
	inputCh := startScan()
	wg.Add(1)

	c := *cl
	go func() {
		for {
			in := <-inputCh
			switch in {
			case "q", "quit":
				wg.Done()
				c.Stop()
			case "h":
				printHelp()
			case "p", "print", "show":
				c.Show()
			}
			fmt.Println("")
		}
	}()
	wg.Wait()
}

func printHelp() {
	fmt.Println("commands: \n\th: help\n\tq: quit\n\tp: print")
}

func startScan() chan string {
	ch := make(chan string)

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			s := scanner.Text()
			if s != "" {
				ch <- s
			}
		}
	}()
	return ch
}

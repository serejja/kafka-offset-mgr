package main

import (
	"fmt"
	"github.com/serejja/kafka-offset-mgr/manager"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "kafka-offset-mgr"
	app.HelpName = "kafka-offset-mgr"
	app.Usage = "Apache Kafka consumer offset manager"
	app.UsageText = "kafka-offset-mgr command [command options] [arguments...]"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:  "view",
			Usage: "View consumer offsets",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "brokers, b",
					Usage: "Comma separated list of Kafka brokers to get offsets from in <host>:<port> form.",
				},
				cli.StringFlag{
					Name:  "groups, g",
					Usage: "Comma separated list of Kafka consumer groups to get offsets for. Optional.",
				},
				cli.StringFlag{
					Name:  "topics, t",
					Usage: "Comma separated list of Kafka topics to get offsets for. Optional.",
				},
			},
			Action: manager.ViewAction,
		},
		{
			Name:  "set",
			Usage: "Set consumer offsets",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "brokers, b",
					Usage:       "Comma separated list of Kafka brokers in `<host>:<port>` form.",
					Destination: &manager.SetBrokers,
				},
				cli.StringFlag{
					Name:        "group, g",
					Usage:       "Kafka consumer group to set offset for.",
					Destination: &manager.SetGroup,
				},
				cli.StringFlag{
					Name:        "topic, t",
					Usage:       "Kafka topic to set offset for.",
					Destination: &manager.SetTopic,
				},
				cli.IntFlag{
					Name:        "partition, p",
					Usage:       "Kafka topic partition to set offset for.",
					Destination: &manager.SetPartition,
				},
				cli.IntFlag{
					Name:        "offset, o",
					Usage:       "New offset value.",
					Destination: &manager.SetOffset,
				},
			},
			Action: manager.SetAction,
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		err := cli.ShowAppHelp(c)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

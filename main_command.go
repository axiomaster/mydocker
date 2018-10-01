package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"mydocker/container"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `create a container`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(ctx cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := ctx.Args().Get(0)
		tty := ctx.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init contanier process",
	Action: func(ctx cli.Context) error {
		log.Infof("init come on")
		cmd := ctx.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}

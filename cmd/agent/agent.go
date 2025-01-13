package main

import (
	"context"
	"os"

	"github.com/lucheng0127/bmsVpcGateway/pkg/agent"
	"github.com/lucheng0127/bmsVpcGateway/pkg/signals"
	"github.com/urfave/cli/v3"
	"k8s.io/klog/v2"
)

func run(ctx context.Context, cmd *cli.Command) error {
	// Init agent
	agent, err := agent.NewAgent(cmd.String("kconfig"))
	if err != nil {
		return err
	}

	// Launch
	return agent.Controller.Run(ctx)
}

func main() {
	cmd := &cli.Command{
		Name:   "agent",
		Action: run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "kconfig",
				Required: false,
				Usage:    "k8s config file path",
			},
		},
	}

	ctx := signals.SetupSignalHandler()
	if err := cmd.Run(ctx, os.Args); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}

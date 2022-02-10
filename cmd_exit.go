package main

import (
	"github.com/urfave/cli/v2"
)

func exitCmd() *cli.Command {
	return &cli.Command{
		Name:      "exit",
		Usage:     "exit neno cli",
		UsageText: "退出neno",
		Action: func(c *cli.Context) error {

			return cli.Exit("玩的开心", 0)

		},
	}
}

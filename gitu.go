package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-util"
	app.Usage = "Command Line Tool for Git"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "userinfo",
			Usage: "show user information.",
			Action: func(c *cli.Context) {
				var (
					cmdOut []byte
					err    error
				)
				cmdName := "git"
				cmdArgs := []string{"config", "--list"}
				if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
					fmt.Fprintln(os.Stderr, "There was an error running git log command: ", err)
					os.Exit(1)
				}
				sha := string(cmdOut)
				fmt.Println(sha)
			},
		},
		{
			Name:  "setuser",
			Usage: "set user information.",
			Action: func(c *cli.Context) {
				var (
					cmdOut []byte
					err    error
				)
				cmdName := "git"
				cmdArgs := []string{"config", "--list"}
				if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
					fmt.Fprintln(os.Stderr, "There was an error running git log command: ", err)
					os.Exit(1)
				}
				sha := string(cmdOut)
				fmt.Println(sha)
			},
		},
	}

	app.Run(os.Args)
}

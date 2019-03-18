package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/namahu/bissucket/config"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	app := cli.NewApp()
	app.Name = "bissucket"
	app.HelpName = "bissucket"
	app.Version = "0.1.1"
	app.Usage = "bissucket is a tool to manipulate Bitbucket Issue from the CLI."
	app.UsageText = "bissucket [global options] command [command options] [arguments...]"

	setFlag := cli.StringFlag{
		Name:  "set, s",
		Usage: "Flag for registering or changing settings.",
	}
	detailFlag := cli.IntFlag{
		Name:  "detail, d",
		Usage: "Specify the issue ID to check details.",
	}
	// titleFlag := cli.StringFlag{
	// 	Name:  "title, t",
	// 	Usage: "Title of Issue.",
	// }
	// priorityFlag := cli.StringFlag{
	// 	Name:  "priority, p",
	// 	Usage: "Priority of Issue.",
	// }
	// kindFlag := cli.StringFlag{
	// 	Name:  "kind, k",
	// 	Usage: "Kind of Issue",
	// }
	// rawcontentFlag := cli.StringFlag{
	// 	Name:  "raw-content, raw",
	// 	Usage: "content of Issue",
	// }

	// コンフィグファイルのチェック。なければ作成
	app.Before = func(c *cli.Context) error {

		var bitbucketUserName string
		var bitbucketPassword string

		if err := config.CheckConfig(); err != nil {
			fmt.Println("Error: No configfile was found. \nWe will start initial setting from now.")
			fmt.Println("")

			fmt.Print("Please enter the password of Bitbucket: ")

			pass, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return fmt.Errorf("ReadPasswordError: %s", err)
			}

			bitbucketPassword = string(pass)

			fmt.Println("")
			fmt.Print("Please enter the user name of Bitbucket: ")
			fmt.Scan(&bitbucketUserName)

			if err = config.CreateConfigFile(bitbucketUserName, bitbucketPassword); err != nil {
				return fmt.Errorf("Error: %s", err)
			}

			fmt.Println("")
			fmt.Println("Creation of config file succeeded.")
			fmt.Println("")
			fmt.Println("Enter the following command for Bitbucket's Synchronize the repository.")
			fmt.Println("")
			fmt.Println("bissucket sync")
			fmt.Println("")

			os.Exit(0)

		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:      "list",
			Usage:     "Display issues for the specified repository.",
			UsageText: "bissucket list [REPOSITORY NAME]",
			Action:    List,
			Flags: []cli.Flag{
				detailFlag,
			},
		},
		{
			Name:      "config",
			Usage:     "Command to set bissucket related operations. If there is no argument, display a list of settings.",
			UsageText: "bissucket config",
			Action:    Config,
			Flags: []cli.Flag{
				setFlag,
			},
		},
		{
			Name:      "show",
			Usage:     "Display issues detail for the specified repository.",
			UsageText: "bissucket show [repository name] [issue id]",
			Action:    Show,
		},
		// {
		// 	Name: "create",
		// 	Usage: "Create the new issue.",
		// 	UsageText: "bissucket create "
		// },
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

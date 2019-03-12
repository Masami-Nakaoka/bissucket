package main

import (
	"fmt"
	"os"
	"syscall"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	app := cli.NewApp()
	app.Name = "bissucket"
	app.HelpName = "bissucket"
	app.Version = "0.1.1"
	app.Usage = "bissucket is a tool to manipulate Bitbucket Issue from the CLI.\n    First from [bissucket sync] please."
	app.UsageText = "bissucket [global options] command [command options] [arguments...]"

	repoFlag := cli.BoolFlag{
		Name:  "repository, r",
		Usage: "Flag for declaring the operation of the repository.",
	}
	issueFlag := cli.StringFlag{
		Name:  "issue, i",
		Usage: "Flag for declaring the operation of the issue.",
	}
	saveCacheFlag := cli.BoolFlag{
		Name: "save-cache, save",
		Usage: "Flag to save issues in cache file."
	}
	setFlag := cli.StringFlag{
		Name:  "set, s",
		Usage: "Flag for registering or changing settings.",
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

		bitbucketUserName = config.GetConfigValueByKey("bitbucketUserName")
		bitbucketPassword = config.GetConfigValueByKey("bitbucketPassword")

		app.Metadata = map[string]interface{}{
			"bitbucketUserName": bitbucketUserName,
			"bitbucketPassword": bitbucketPassword,
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:      "sync",
			Usage:     "Synchronize with Bitbucket's repository and issue.",
			UsageText: "bissucket sync [--repository, -r][--issue, -i]",
			Action:    Sync,
			Flags: []cli.Flag{
				repoFlag,
				issueFlag,
			},
		},
		{
			Name:      "list",
			Usage:     "Display Issue and list of repositories. Display a list of defaultRepository if no options are given.",
			UsageText: "bissucket list [--repository, -r] [--issue REPOSITORY NAME, -i REPOSITOORY NAME] [--save-cache]",
			Action:    List,
			Flags: []cli.Flag{
				repoFlag,
				issueFlag,
				saveCacheFlag,
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
			Usage:     "Display Issue details of defaultRepository.",
			UsageText: "bissucket show [issue id]",
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

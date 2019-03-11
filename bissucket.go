package main

import (
	"fmt"
	"os"
	"syscall"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	"bitbucket.org/Masami_Nakaoka/bissucket/issue"
	repo "bitbucket.org/Masami_Nakaoka/bissucket/repository"
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

	listFlag := cli.BoolFlag{
		Name:  "list, l",
		Usage: "Show your repository list.",
	}
	repoFlag := cli.BoolFlag{
		Name:  "repository, r",
		Usage: "Flag for declaring the operation of the repository.",
	}
	repoNameFlag := cli.StringFlag{
		Name:  "repository-name, rn",
		Usage: "Flag when specifying `repository name`",
	}
	issueFlag := cli.BoolFlag{
		Name:  "issue, i",
		Usage: "Flag for declaring the operation of the issue.",
	}
	setFlag := cli.StringFlag{
		Name:  "set, s",
		Usage: "Flag for registering or changing settings.",
	}
	// detailFlag := cli.IntFlag{
	// 	Name:  "detail, d",
	// 	Usage: "Display issue details.",
	// }
	titleFlag := cli.StringFlag{
		Name:  "title, t",
		Usage: "Title of Issue.",
	}
	priorityFlag := cli.StringFlag{
		Name:  "priority, p",
		Usage: "Priority of Issue.",
	}
	kindFlag := cli.StringFlag{
		Name:  "kind, k",
		Usage: "Kind of Issue",
	}
	rawcontentFlag := cli.StringFlag{
		Name:  "raw-content, raw",
		Usage: "content of Issue",
	}

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
			UsageText: "bissucket list [--repository, -r] [--repository-name REPOSITORY NAME, -rn REPOSITOORY NAME]",
			Action:    List,
			Flags: []cli.Flag{
				repoFlag,
				repoNameFlag,
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
			Name:      "repository",
			Aliases:   []string{"repo"},
			Usage:     "Display a list of repositories or set a default repository.",
			UsageText: "bissucket repository --list",
			Flags: []cli.Flag{
				listFlag,
			},
			Action: repo.RepositoryList,
			Subcommands: []cli.Command{
				{
					Name:      "default-set",
					Aliases:   []string{"df"},
					Usage:     "Set the default Repository.",
					UsageText: "bissucket repository default-set [repository name]",
					Action:    repo.SetDefaultRepository,
				},
			},
		},
		{
			Name:      "issue",
			Aliases:   []string{"i"},
			Usage:     "Command to operate Issue.",
			UsageText: "bissucket issue [command][command options]",
			Subcommands: []cli.Command{
				{
					Name:      "create",
					Aliases:   []string{"c"},
					Usage:     "Create an issue.",
					UsageText: "bissucket issue create [command options]",
					Flags: []cli.Flag{
						titleFlag,
						priorityFlag,
						kindFlag,
						rawcontentFlag,
					},
					Action: issue.Create,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

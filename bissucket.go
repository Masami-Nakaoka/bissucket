package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	issue "bitbucket.org/Masami_Nakaoka/bissucket/issue"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	configPath          = os.Getenv("HOME")
	repositoryCachePath = os.Getenv("HOME") + "/.bissucket.repositoriescache.json"
)

const (
	configFileName = ".bissucket.config"
	configFileType = "json"
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
	// detailFlag := cli.IntFlag{
	// 	Name:  "detail, d",
	// 	Usage: "Display issue details.",
	// }
	repoNameFlag := cli.StringFlag{
		Name:  "repository, r",
		Usage: "Input repository name.",
	}
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
		viper.SetConfigName(configFileName)
		viper.AddConfigPath(configPath)
		viper.AddConfigPath(".")

		var bitbucketUserName string
		var bitbucketPassword string

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error: No configfile was found. We will start initial setting from now.")
			fmt.Println("")

			fmt.Print("Please enter the password of Bitbucket: ")

			pass, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return fmt.Errorf("ReadPasswordError: %s", err)
			}

			bitbucketPassword = string(pass)

			viper.Set("bitbucketPassword", bitbucketPassword)

			fmt.Println("")
			fmt.Print("Please enter the user name of Bitbucket: ")
			fmt.Scan(&bitbucketUserName)
			viper.Set("bitbucketUserName", bitbucketUserName)

			configJSON, err := json.MarshalIndent(viper.AllSettings(), "", "    ")
			if err != nil {
				return fmt.Errorf("JsonMarshalError: %s", err)
			}

			err = ioutil.WriteFile(filepath.Join(configPath, configFileName+"."+configFileType), configJSON, os.ModePerm)
			if err != nil {
				return fmt.Errorf("WriteFileError: %s", err)
			}

			fmt.Println("")
			fmt.Println("Creation of config file succeeded.")
			fmt.Println("Enter the following command for Bitbucket's Synchronize the repository.")
			fmt.Println("")
			fmt.Println("bissucket sync")
			fmt.Println("")

			os.Exit(0)

		}

		app.Metadata = map[string]interface{}{
			"bitbucketUserName": viper.GetString("bitbucketUserName"),
			"bitbucketPassword": viper.GetString("bitbucketPassword"),
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:      "repository",
			Aliases:   []string{"repo"},
			Usage:     "Repository related operations. Currently only list view.",
			UsageText: "bissucket repository --list",
			Flags: []cli.Flag{
				listFlag,
			},
			Action: Repository,
		},
		{
			Name:      "issue",
			Aliases:   []string{"i"},
			Usage:     "Command to operate Issue.",
			UsageText: "bissucket issue [command][command options]",
			Subcommands: []cli.Command{
				{
					Name:      "list",
					Aliases:   []string{"l"},
					Usage:     "Display Issue list of specified Repository",
					UsageText: "bissucket issue list -r [repository name]",
					Flags: []cli.Flag{
						repoNameFlag,
					},
					Action: issue.IssueList,
				},
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
		{
			Name:      "sync",
			Usage:     "Get your repository from Bitbucket.",
			UsageText: "bissucket sync",
			Action:    Sync,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

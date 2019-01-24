package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

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
	bitbucketURI   = "https://api.bitbucket.org/2.0/"
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
	detailFlag := cli.IntFlag{
		Name:  "detail, d",
		Usage: "Display issue details.",
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
			Usage:     "Display the issue of a specific repository.",
			UsageText: "bissucket issue [repository name]",
			Flags: []cli.Flag{
				detailFlag,
			},
			Action: Issue,
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

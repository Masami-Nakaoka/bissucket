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

var configPath = os.Getenv("Home")

const (
	configFileName = ".bissucket.config"
	configFileType = "json"
	bitbucketURI   = "https://api.bitbucket.org/2.0/"
)

func main() {
	fmt.Println(configPath)
	app := cli.NewApp()
	app.Name = "bissucket"
	app.Version = "0.0.1"
	app.Usage = "bissucket is a cli tool to manipulate bitbucket issues"

	// listFlag := cli.BoolFlag{
	// 	Name:  "list, l",
	// 	Usage: "Display data list",
	// }

	syncFlag := cli.BoolFlag{
		Name:  "sync, s",
		Usage: "Get your repository from Bitbucket",
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
			fmt.Print("Please enter the user name of Bitbucket: ")
			fmt.Scan(&bitbucketUserName)

			fmt.Print("Please enter the password of Bitbucket: ")

			pass, err := terminal.ReadPassword(syscall.Stdin)
			if err != nil {
				fmt.Print(err)
			} else {
				bitbucketPassword = string(pass)
			}
			viper.Set("bitbucketUserName", bitbucketUserName)
			viper.Set("bitbucketPassword", bitbucketPassword)

			configJSON, err := json.MarshalIndent(viper.AllSettings(), "", "    ")
			if err != nil {
				return fmt.Errorf("Error: %s", err)
			}

			err = ioutil.WriteFile(filepath.Join(configPath, configFileName+"."+configFileType), configJSON, os.ModePerm)
			if err != nil {
				return fmt.Errorf("Error: %s", err)
			}

			fmt.Println("")
			fmt.Println("Creation of config file succeeded.")
			fmt.Println("Enter the following command for Bitbucket's Synchronize the repository.")
			fmt.Println("")
			fmt.Println("bissucket repository --sync")
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
			Name:    "repository",
			Aliases: []string{"repo"},
			Flags: []cli.Flag{
				syncFlag,
			},
			Action: Repository,
		},
	}

	fmt.Println("Hello")

	app.Run(os.Args)
}

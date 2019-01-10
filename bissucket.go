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
		viper.SetConfigType(configFileType)
		viper.SetConfigFile(configFileName)
		viper.AddConfigPath(configPath)

		var bitbucketUserName string
		var bitbucketToken string

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error: No configfile was found. We will start initial setting from now.")
			fmt.Println("Please enter the user name of Bitbucket.")
			fmt.Scan(&bitbucketUserName)
			viper.Set("bitbucketUserName", bitbucketUserName)

			fmt.Println("")

			fmt.Println("Please enter the token of Bitbucket.")
			fmt.Println("※The entered contents are not displayed.")
			token, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return fmt.Errorf("Error: %s", err)
			}

			bitbucketToken = string(token)
			viper.Set("bitbucketToken", bitbucketToken)

			configJSON, err := json.MarshalIndent(viper.AllSettings(), "", "    ")
			if err != nil {
				return fmt.Errorf("Error: %s", err)
			}

			err = ioutil.WriteFile(filepath.Join(configPath, configFileName+"."+configFileType), configJSON, os.ModePerm)
			if err != nil {
				return fmt.Errorf("Error: %s", err)
			}

		}

		app.Metadata = map[string]interface{}{
			"bitbucketUserName": viper.GetString("bitbucketUserName"),
			"bitbucketToken":    viper.GetString("bitbucketToken"),
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
			Action: func(c *cli.Context) error {
				if c.Bool("sync") {
					fmt.Println("Sync")
				} else {
					fmt.Println("Not sync")
				}

				return nil
			},
		},
	}

	fmt.Println("Hello")

	app.Run(os.Args)
}

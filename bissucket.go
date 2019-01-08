package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/urfave/cli"
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

	// コンフィグファイルのチェック。なければ作成
	app.Before = func(c *cli.Context) error {
		viper.SetConfigType(configFileType)
		viper.SetConfigFile(configFileName)
		viper.AddConfigPath(configPath)

		var bitbucketUserName string
		var bitbucketPassWord string

		if err := viper.ReadConfig(); err != nil {
			fmt.Println("Error: No configfile was found. We will start initial setting from now")
			fmt.Println("Please enter the user name of Bitbucket")
			fmt.Scan(&bitbucketUserName)
			viper.Set("bitbucketUserName", bitbucketUserName)

			fmt.Println("Please enter the password of Bitbucket")
			fmt.Scan(&bitbucketPassWord)
			viper.Set("bitbucketPassWord", bitbucketPassWord)
		}
	}

	fmt.Println("Hello")

	app.Run(os.Args)
}

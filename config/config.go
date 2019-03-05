// package config

// import (
// 	"os"
// 	"github.com/spf13/viper"
// )

// var (
// 	configPath          = os.Getenv("HOME")
// 	repositoryCachePath = os.Getenv("HOME") + "/.bissucket.repositoriescache.json"
// )

// const (
// 	configFileName = ".bissucket.config"
// 	configFileType = "json"
// )

// func CheckConfig() {
// 	viper.SetConfigName(configFileName)
// 	viper.AddConfigPath(configPath)
// 	viper.AddConfigPath(".")

// 	if err != viper.ReadConfig(); err != nil {
// 		fmt.Println("Error: No configfile was found. We will start initial setting from now.")
// 		fmt.Println("")

// 		fmt.Print("Please enter the password of Bitbucket: ")

// 		pass, err := terminal.ReadPassword(int(syscall.Stdin))
// 		if err != nil {
// 			return fmt.Errorf("ReadPasswordError: %s", err)
// 		}

// 		bitbucketPassword = string(pass)

// 		viper.Set("bitbucketPassword", bitbucketPassword)

// 		fmt.Println("")
// 		fmt.Print("Please enter the user name of Bitbucket: ")
// 		fmt.Scan(&bitbucketUserName)
// 		viper.Set("bitbucketUserName", bitbucketUserName)

// 	}

// }

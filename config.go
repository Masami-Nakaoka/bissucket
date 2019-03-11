package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	"github.com/urfave/cli"
)

func Config(c *cli.Context) error {

	if c.String("s") == "" {

		var configs map[string]string

		buf, err := config.GetAllConfigKeyAndValue()
		if err != nil {
			return err
		}

		if err = json.Unmarshal(buf, &configs); err != nil {
			return fmt.Errorf("UnmarshallError: %s", err)
		}

		fmt.Println("")

		for key, value := range configs {
			fmt.Printf("%s=%s\n", key, value)
		}

		fmt.Println("")

	} else {

		configStrSlice := strings.Split(c.String("s"), "=")
		configKey := configStrSlice[0]
		configValue := configStrSlice[1]

		if err := config.SetConfigKeyAndValue(configKey, configValue); err != nil {
			return fmt.Errorf("SetConfigError: %s", err)
		}

		fmt.Println("The setting was updated.")
	}

	return nil
}

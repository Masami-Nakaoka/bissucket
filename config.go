package main

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	"github.com/urfave/cli"
)

var configs map[string]string

func Config(c *cli.Context) error {

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

	return nil
}

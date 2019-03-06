package repository

import (
	"errors"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	"github.com/urfave/cli"
)

func SetDefaultRepository(c *cli.Context) error {

	if c.NArg() > 1 {

		return errors.New("Please input only one argument")

	}
	if c.Args().First() == "" {

		return errors.New("Please enter one repository name")

	}

	configKey := "defaultRepository"
	defaultRepositoryName := c.Args().First()

	config.SetConfigKeyAndValue(configKey, defaultRepositoryName)

	return nil
}

package main

import (
	"errors"

	"github.com/urfave/cli"
)

func Create(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("The title of the issue is a required item.")
	}

	return nil
}

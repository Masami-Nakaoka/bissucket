package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func Repository(c *cli.Context) error {
	userName := c.App.Metadata["bitbucketUserName"]
	fmt.Println(userName)
	return nil
}

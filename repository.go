package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func Repository(c *cli.Context) error {
	userName := c.App.Metadata["bitbucketUserName"]
	token := c.App.Metadata["bitbucketToken"]
	fmt.Println(userName)
	fmt.Println(token)
	return nil
}

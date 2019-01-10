package main

import (
	"fmt"

	"github.com/urfave/cli"
)

const (
    baseUrl = "https//api.bitbucket.org/2.0/repositories/"
)

func Repository(c *cli.Context) error {
	userName := c.App.Metadata["bitbucketUserName"]
	token := c.App.Metadata["bitbucketToken"]
	
	endPoint := baseUrl + userName + "/"
	
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return err
	}
	return nil
}

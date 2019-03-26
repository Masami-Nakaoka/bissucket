package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/namahu/bissucket/config"
	bitbucket "github.com/namahu/bissucket/lib"
	"github.com/urfave/cli"
)

func Create(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("The title of the issue is a required item.")
	}

	userName := config.GetConfigValueByKey("bitbucketUserName")
	repoName := c.Args()[0]

	issue := url.Values{}
	issue.Add("title", c.Args()[1])

	endPoint := "repositories/" + userName + "/" + repoName + "/issues"

	res, err := bitbucket.DoPost(endPoint, userName, issue)
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

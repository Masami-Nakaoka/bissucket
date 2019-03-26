package main

import (
	"errors"
	"fmt"

	"github.com/namahu/bissucket/config"
	bitbucket "github.com/namahu/bissucket/lib"
	"github.com/urfave/cli"
)

var (
	listItem = map[string][]string{
		"kind":     {"bug", "enhancement", "proposal", "task"},
		"priority": {"trivial", "minor", "major", "critical", "blocker"},
	}
)

func existenceCheck(listName string, list []string, checkValue string) error {
	for _, v := range list {
		if v == checkValue {
			return nil
		}
	}
	return fmt.Errorf("Specify one of the following for %v: %v", listName, list)
}

func Create(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("The title of the issue is a required item.")
	}

	userName := config.GetConfigValueByKey("bitbucketUserName")
	repoName := c.Args()[0]

	issue := make(map[string]interface{})
	issue["title"] = c.Args()[1]
	if c.String("priority") != "" {
		err := existenceCheck("priority", listItem["priority"], c.String("priority"))
		if err != nil {
			return err
		}
		issue["priority"] = c.String("priority")
	}
	if c.String("kind") != "" {
		err := existenceCheck("kind", listItem["kind"], c.String("kind"))
		if err != nil {
			return err
		}
		issue["kind"] = c.String("kind")
	}
	if c.String("raw-content") != "" {
		contentMap := map[string]string{
			"raw": c.String("raw-content"),
		}
		issue["content"] = contentMap
	}

	endPoint := "repositories/" + userName + "/" + repoName + "/issues"

	err := bitbucket.DoPost(endPoint, userName, issue)
	if err != nil {
		return err
	}

	fmt.Println("Issue create sucess")

	return nil
}

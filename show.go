package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	bitbucket "bitbucket.org/Masami_Nakaoka/bissucket/lib"
	"github.com/urfave/cli"
)

func showIssueDetail(issues *bitbucket.Issues, issueID string, repositoryName string) error {

	fmt.Println(issues)

	fmt.Println("------------------------------")
	fmt.Println("Details of " + repositoryName + "'s issue")
	fmt.Println("------------------------------")

	for _, issue := range issues.Values {
		if issueID == strconv.Itoa(issue.ID) {
			fmt.Print("ID:        ")
			fmt.Println(issue.ID)
			fmt.Println("State:    " + issue.State)
			fmt.Println("Priority: " + issue.Priority)
			fmt.Println("Kind:     " + issue.Kind)
			fmt.Println("Assignee: " + issue.Assignee.Username)
			fmt.Println("Title:    " + issue.Title)
			fmt.Println("Content:  ")
			fmt.Println(issue.Content.Raw)
			fmt.Println("")
		}
	}

	return nil
}

func Show(c *cli.Context) error {

	if c.NArg() > 1 {
		return errors.New("Please give one argument.")
	}

	issueID := c.Args().First()
	repositoryName := config.GetConfigValueByKey("defaultRepository")

	buf, err := getListByCache("issue")
	if err != nil {
		return fmt.Errorf("CacheReadError: %s", err)
	}

	if err = json.Unmarshal(buf, &issues); err != nil {
		return fmt.Errorf("JsonUnmarshalError: %s", err)
	}

	showIssueDetail(issues, issueID, repositoryName)
	fmt.Println(issueID)
	return nil
}

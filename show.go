package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	bitbucket "bitbucket.org/Masami_Nakaoka/bissucket/lib"
	"github.com/urfave/cli"
)

func showIssueDetail(issues *bitbucket.Issues, issueID string) error {

	fmt.Println("------------------------------")
	fmt.Println("Details of issue")
	fmt.Println("------------------------------")

	for _, issue := range issues.Values {
		if issueID == strconv.Itoa(issue.ID) {
			fmt.Printf("ID:         %d\n", issue.ID)
			fmt.Printf("Repository: %s\n", issue.Repository.Name)
			fmt.Printf("State:      %s\n", issue.State)
			fmt.Printf("Priority:   %s\n", issue.Priority)
			fmt.Printf("Kind:       %s\n", issue.Kind)
			fmt.Printf("Assignee:   %s\n", issue.Assignee.Username)
			fmt.Printf("Title:      %s\n", issue.Title)
			fmt.Printf("Content:    %s\n", issue.Content.Raw)
		}
	}

	return nil
}

func Show(c *cli.Context) error {

	if c.NArg() > 1 {
		return errors.New("Please give one argument.")
	}

	issueID := c.Args().First()

	buf, err := getListByCache("issue")
	if err != nil {
		return fmt.Errorf("CacheReadError: %s", err)
	}

	if err = json.Unmarshal(buf, &issues); err != nil {
		return fmt.Errorf("JsonUnmarshalError: %s", err)
	}

	showIssueDetail(issues, issueID)

	return nil
}

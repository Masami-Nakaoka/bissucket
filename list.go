package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/namahu/bissucket/config"
	bitbucket "github.com/namahu/bissucket/lib"
	"github.com/urfave/cli"
)

var issues *bitbucket.Issues

func showIssueList(issues *bitbucket.Issues) {

	issueItemList := make([][]string, 0, len(issues.Values))
	for _, issue := range issues.Values {
		issueItemList = append(issueItemList, []string{
			strconv.Itoa(issue.ID),
			issue.Repository.Name,
			issue.State,
			issue.Priority,
			issue.Kind,
			issue.Assignee.Username,
			issue.Title,
		})
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 2, ' ', 0)

	defer w.Flush()

	header := []string{"ID", "Repo", "State", "Priority", "Kind", "Assignee", "Title"}
	fmt.Fprintln(w, strings.Join(header[:], "\t"))

	for _, issueItem := range issueItemList {
		strings := strings.Join(issueItem[:], "\t")
		fmt.Fprintln(w, strings)
	}
}

func List(c *cli.Context) error {

	if c.NArg() == 0 && c.Args().First() == "" {

		return errors.New("Enter repository name to display issues.")

	}

	repositoryName := c.Args().First()
	userName := config.GetConfigValueByKey("bitbucketUserName")
	endPoint := "repositories/" + userName + "/" + repositoryName + "/issues"

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("fecthError: %s", err)
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&issues); err != nil {
		return fmt.Errorf("jsonDecodeError: %s", err)
	}

	showIssueList(issues)

	return nil
}

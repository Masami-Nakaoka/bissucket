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

var issues interface{}

func showIssuesList(issues *bitbucket.Issues) {
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

func showIssueList(issue *bitbucket.Issue) {
	itemList := [][]string{
		[]string{"IssueID", strconv.Itoa(issue.ID)},
		[]string{"Repository", issue.Repository.Name},
		[]string{"State", issue.State},
		[]string{"Priority", issue.Priority},
		[]string{"Kind", issue.Kind},
		[]string{"Assignee", issue.Assignee.Username},
		[]string{"Title", issue.Title},
		[]string{"Content", "\n" + issue.Content.Raw},
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 2, ' ', 0)

	defer w.Flush()

	for _, item := range itemList {
		strings := strings.Join(item[:], "\t")
		fmt.Fprintln(w, strings)
	}
}

func List(c *cli.Context) error {

	if c.NArg() == 0 && c.Args().First() == "" {

		return errors.New("Enter repository name to display issues.")

	}

	var (
		repositoryName = c.Args().First()
		userName       = config.GetConfigValueByKey("bitbucketUserName")
		endPoint       = "repositories/" + userName + "/" + repositoryName + "/issues"
	)

	if c.Int("d") == 0 {
		issues = &bitbucket.Issues{}
	} else if c.Int("d") != 0 {
		issueID := c.Int("d")
		endPoint += "/" + strconv.Itoa(issueID)
		issues = &bitbucket.Issue{}
	}

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("fecthError: %s", err)
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&issues); err != nil {
		return fmt.Errorf("jsonDecodeError: %s", err)
	}

	switch i := issues.(type) {
	case *bitbucket.Issues:
		showIssuesList(i)
	case *bitbucket.Issue:
		showIssueList(i)
	default:
		return fmt.Errorf("Unexpected type: %T", i)
	}

	return nil
}

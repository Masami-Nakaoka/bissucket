package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/namahu/bissucket/config"
	bitbucket "github.com/namahu/bissucket/lib"
	"github.com/urfave/cli"
)

type IssueCache struct {
	Repository string
	Store      *bitbucket.Issues
}

func saveIssuesInCache(issueCache *IssueCache) error {

	issueCachePath := config.GetConfigValueByKey("issueCachePath")

	buf, err := json.MarshalIndent(issueCache, "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshallError: %s", err)
	}

	err = ioutil.WriteFile(issueCachePath, buf, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	return nil
}

func showIssueList(issueCache *IssueCache) {

	issueItemList := make([][]string, 0, len(issueCache.Store.Values))
	for _, issue := range issueCache.Store.Values {
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

func getIssueList(repoName string, cache *IssueCache) error {

	userName := config.GetConfigValueByKey("bitbucketUserName")
	endPoint := "repositories/" + userName + "/" + repoName + "/issues"

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("FetchError: %s", err)
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&cache.Store)

}

func List(c *cli.Context) error {

	issueCache := &IssueCache{}

	if c.NArg() == 0 && c.Args().First() == "" {
		return errors.New("Enter repository name to display issues.")
	}

	repositoryName := c.Args().First()
	issueCache.Repository = repositoryName

	if err := loadCache(issueCache); err != nil {
		return fmt.Errorf("loadCacheError: %s", err)
	}

	if repositoryName != issueCache.Repository {
		if err := getIssueList(repositoryName, issueCache); err != nil {
			return fmt.Errorf("getIssueListError: %s", err)
		}
		issueCache.Repository = repositoryName
	}

	if c.Bool("save") {

		saveIssuesInCache(issueCache)
	}

	showIssueList(issueCache)

	return nil
}

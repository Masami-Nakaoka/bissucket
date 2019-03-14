package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

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

func showIssueList(repositoryName string, store *bitbucket.Issues) {

	fmt.Println("------------------------------")
	fmt.Println("Issue List of " + repositoryName)
	fmt.Println("------------------------------")
	fmt.Println("ID / State / Priority / Kind / Assignee / Title")

	var issueTemplate string
	for _, issue := range store.Values {
		issueTemplate = strconv.Itoa(issue.ID) + " / " + issue.State + " / " + issue.Priority + " / " + issue.Kind + " / " + issue.Assignee.Username + " / " + issue.Title
		fmt.Println(issueTemplate)
	}
}

func getListByCache(target string) ([]byte, error) {
	cachePath := config.GetConfigValueByKey(target + "CachePath")

	return ioutil.ReadFile(cachePath)

}

func getIssueList(repoName string, store *bitbucket.Issues) error {

	userName := config.GetConfigValueByKey("bitbucketUserName")
	endPoint := "repositories/" + userName + "/" + repoName + "/issues"

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("FetchError: %s", err)
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&store)

}

func List(c *cli.Context) error {

	var store bitbucket.Issues

	if c.NArg() == 0 && c.Args().First() == "" {
		return fmt.Errorf("Enter repository name to display issues.")
	}

	repositoryName := c.Args().First()

	if err := getIssueList(repositoryName, &store); err != nil {
		return fmt.Errorf("getIssueListError: %s", err)
	}

	if c.Bool("save") {

		issueCache := &IssueCache{
			Repository: repositoryName,
			Store:      &store,
		}

		saveIssuesInCache(issueCache)
	}

	showIssueList(repositoryName, &store)

	return nil
}

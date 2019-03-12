package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	bitbucket "bitbucket.org/Masami_Nakaoka/bissucket/lib"
	"github.com/urfave/cli"
)

func showIssueList(repositoryName string, issues *bitbucket.Issues) {

	fmt.Println("------------------------------")
	fmt.Println("Issue List of " + repositoryName)
	fmt.Println("------------------------------")
	fmt.Println("ID / State / Priority / Kind / Assignee / Title")

	var issueTemplate string
	for _, issue := range issues.Values {
		issueTemplate = strconv.Itoa(issue.ID) + " / " + issue.State + " / " + issue.Priority + " / " + issue.Kind + " / " + issue.Assignee.Username + " / " + issue.Title
		fmt.Println(issueTemplate)
	}
}

func showRepositoryList(repos *bitbucket.Repos) {

	fmt.Println("---------------------------------")
	fmt.Println("Repository Name  /  Has Issues")
	fmt.Println("---------------------------------")

	for _, repo := range repos.Values {
		fmt.Printf("%s  /  %t\n", repo.Name, repo.HasIssues)
	}

	fmt.Println("")

}

func getListByCache(target string) ([]byte, error) {
	cachePath := config.GetConfigValueByKey(target + "CachePath")

	return ioutil.ReadFile(cachePath)

}

func getIssueListByRepositoryName(repoName string) ([]byte, error) {

	userName := config.GetConfigValueByKey("bitbucketUserName")
	endPoint := "repositories/" + userName + "/" + repoName + "/issues"

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return nil, fmt.Errorf("FetchError: %s", err)
	}

	defer res.Body.Close()

	var issues *bitbucket.Issues

	if err = json.NewDecoder(res.Body).Decode(&issues); err != nil {
		return nil, fmt.Errorf("JsonDecodeError: %s", err)
	}

	return json.MarshalIndent(issues, "", "    ")

}

func List(c *cli.Context) error {

	userName := config.GetConfigValueByKey("bitbucketUserName")
	endPoint := "repositories/" + userName

	if c.Bool("r") {

		endPoint += "?pagelen=100"

		res, err := bitbucket.DoGet(endPoint, userName)
		if err != nil {
			return fmt.Errorf("fetchError: %s", err)
		}

		defer res.Body.Close()

		if err = json.NewDecoder(res.Body).Decode(&repos); err != nil {
			return fmt.Errorf("jsonDecodeError: %s", err)
		}

		showRepositoryList(repos)

	} else if c.String("i") != "" {

		repositoryName := c.String("i")
		endPoint += "/" + repositoryName + "/issues"

		res, err := bitbucket.DoGet(endPoint, userName)
		if err != nil {
			return fmt.Errorf("fetchError: %s", err)
		}

		defer res.Body.Close()

		if err = json.NewDecoder(res.Body).Decode(&issues); err != nil {
			return fmt.Errorf("jsonDecodeError: %s", err)
		}

		showIssueList(repositoryName, issues)

	}

	return nil
}

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

	fmt.Println("-----------------------\n  Repository Name  \n-----------------------")

	for _, repo := range repos.Values {
		fmt.Printf("%s\n", repo.Name)
	}

	fmt.Println("")

}

func getListByCache(target string) ([]byte, error) {
	cachePath := config.GetConfigValueByKey(target + "CachePath")

	return ioutil.ReadFile(cachePath)

}

func List(c *cli.Context) error {

	var (
		buf []byte
		err error
	)

	if c.Bool("r") {

		buf, err = getListByCache("repository")
		if err != nil {
			return err
		}

		if err = json.Unmarshal(buf, &repos); err != nil {
			return fmt.Errorf("jsonUnMarshallErrror: %s", err)
		}

		showRepositoryList(repos)

	} else {

		repositoryName := config.GetConfigValueByKey("defaultRepository")

		if c.String("rn") != "" {

			// repositoryName := c.String("rn")
		} else {

			buf, err = getListByCache("issue")
			if err != nil {
				return err
			}

		}

		if err = json.Unmarshal(buf, &issues); err != nil {
			return fmt.Errorf("jsonUnMarshallErrror: %s", err)
		}

		showIssueList(repositoryName, issues)

	}

	return nil
}

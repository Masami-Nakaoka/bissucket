package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	bitbucket "bitbucket.org/Masami_Nakaoka/bissucket/lib"
	"github.com/urfave/cli"
)

const homePath = os.Getenv("HOME")

var (
	issueCachePath      = homePath + "/.bissucket.issuecache.json"
	repositoryCachePath = homePath + "/.bitbucket.repositorycache.json"
)

func saveIssuesInCache(issue *bitbucket.Issues) error {

	buf, err := json.MarshalIndent(issue, "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshallError: $s", err)
	}

	err = ioutil.WriteFile(issueCachePath, issue, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	return nil
}

func fetchIssuesByDefaultRepository() error {

	userName := config.GetConfigValueByKey("bitbucketUserName")
	endPoint := "repositories/" + userName + "/" + repositoryName + "/issues"

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("fetchError: %s", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Docode(&bitbucket.Issues)
	if err != nil {
		return fmt.Errorf("DecodeError: %s", err)
	}

	err = saveIssuesInCache(bitbucket.Issues)
	if err != nil {
		return fmt.Errorf("CacheSaveError: %s", err)
	}

	return nil
}

func Sync(c *cli.Context) error {

	if c.Bool("i") {

		fetchIssuesByDefaultRepository()

		return nil

	}

	return nil
}

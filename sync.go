package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	bitbucket "bitbucket.org/Masami_Nakaoka/bissucket/lib"
	"github.com/urfave/cli"
)

var (
	issues *bitbucket.Issues
	repos  *bitbucket.Repos
)

func saveIssuesInCache(issue *bitbucket.Issues) error {

	issueCachePath := config.GetConfigValueByKey("issueCachePath")

	buf, err := json.MarshalIndent(issue, "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshallError: %s", err)
	}

	err = ioutil.WriteFile(issueCachePath, buf, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	return nil
}

func saveRepositoryInCache(r *bitbucket.Repos) error {

	repositoryCachePath := config.GetConfigValueByKey("repositoryCachePath")

	buf, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshallError: $s", err)
	}

	err = ioutil.WriteFile(repositoryCachePath, buf, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	return nil
}

func fetchIssuesByDefaultRepository() error {

	userName := config.GetConfigValueByKey("bitbucketUserName")
	defaultRepository := config.GetConfigValueByKey("defaultRepository")
	endPoint := "repositories/" + userName + "/" + defaultRepository + "/issues"

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("fetchError: %s", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		return fmt.Errorf("DecodeError: %s", err)
	}

	err = saveIssuesInCache(issues)
	if err != nil {
		return fmt.Errorf("CacheSaveError: %s", err)
	}

	return nil
}

func fetchAllRepository() error {

	userName := config.GetConfigValueByKey("bitbucketUserName")
	endPoint := "repositories/" + userName + "?pagelen=100"

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("fetchError: %s", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&repos)
	if err != nil {
		return fmt.Errorf("DecodeError: %s", err)
	}

	err = saveRepositoryInCache(repos)
	if err != nil {
		return fmt.Errorf("CacheSaveError: %s", err)
	}

	return nil
}

func Sync(c *cli.Context) error {

	if c.Bool("i") {

		err := fetchIssuesByDefaultRepository()
		if err != nil {
			return fmt.Errorf("FetchError: %s", err)
		}

	} else if c.Bool("r") {

		err := fetchAllRepository()
		if err != nil {
			return fmt.Errorf("FetchError: %s", err)
		}

	} else if !c.Bool("i") && !c.Bool("r") {
		return errors.New("Specify one option and execute.")
	}

	return nil
}

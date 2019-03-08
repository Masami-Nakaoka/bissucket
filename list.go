package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	"github.com/urfave/cli"
)

func showList(target string) error {
	return nil
}

func getIssueByDefauleRepository() ([]byte, error) {
	issueCachePath := config.GetConfigValueByKey("issueCachePath")

	return ioutil.ReadFile(issueCachePath)

}

func List(c *cli.Context) error {

	var buf []byte

	listTarget := "issue"

	if c.String("rn") {

		// repositoryName := c.String("rn")
	} else {

		buf, err := getIssueByDefauleRepository()
		if err != nil {
			return err
		}

	}

	if err = json.Unmarshal(buf, &issues); err != nil {
		return fmt.Errorf("jsonUnMarshallErrror: %s", err)
	}

	showList(listTarget)

	return nil
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func Sync(c *cli.Context) error {
	userName := c.App.Metadata["bitbucketUserName"].(string)
	pass := c.App.Metadata["bitbucketPassword"].(string)

	endPoint := bitbucketURI + "repositories/" + userName + "?pagelen=100"

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return fmt.Errorf("RequestError: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(userName, pass)

	fmt.Println("Request Set")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("ResponceError: %s", err)
	}

	if res.StatusCode != 200 {
		return errors.New(res.Status)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&repositories)
	if err != nil {
		return fmt.Errorf("JsonDecodeError: %s", err)
	}

	err = saveRepositoryInCache(repositoryCachePath, repositories)
	if err != nil {
		return fmt.Errorf("saveRepositoryError: %s", err)
	}

	return nil
}

func saveRepositoryInCache(filename string, r *Repo) error {
	buf, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshallError: %s", err)
	}

	err = ioutil.WriteFile(filename, buf, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	fmt.Println("Sync sucess")

	return nil

}

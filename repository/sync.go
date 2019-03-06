package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	bitbucket "bitbucket.org/Masami_Nakaoka/bissucket/lib"
	"github.com/urfave/cli"
)

func Sync(c *cli.Context) error {
	userName := c.App.Metadata["bitbucketUserName"].(string)
	pass := c.App.Metadata["bitbucketPassword"].(string)

	endPoint := "repositories/" + userName + "?pagelen=100"

	fmt.Println("Start synchronization")

	res, err := bitbucket.DoGet(endPoint, userName, pass)
	if err != nil {
		return fmt.Errorf("APIError: %s", err)
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

	fmt.Println("Synchronization succeeded")

	return nil
}

func saveRepositoryInCache(filename string, r *Repos) error {
	buf, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshallError: %s", err)
	}

	err = ioutil.WriteFile(filename, buf, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	return nil

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli"
)

func Repository(c *cli.Context) error {
	if c.Bool("l") {
		err := showRepositoryList(repositoryCachePath)
		if err != nil {
			return fmt.Errorf("showRepositoryListError: %s", err)
		}
	}

	return nil
}

func showRepositoryList(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("fileReadError: %s", err)
	}

	err = json.Unmarshal(buf, &repositories)

	fmt.Println("-----------------------\n  Repository Name  \n-----------------------")

	for _, repo := range repositories.Values {
		fmt.Printf("%s\n", repo.Name)
	}

	fmt.Println("")

	return nil
}

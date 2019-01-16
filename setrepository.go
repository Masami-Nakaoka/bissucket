package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/urfave/cli"
)

func SetRepository(c *cli.Context) error {
	if c.NArg() > 1 {
		return errors.New("Too many arguments")
	}

	viper.Set("useRepository", c.Args().Get(0))

	configJSON, err := json.MarshalIndent(viper.AllSettings(), "", "    ")
	if err != nil {
		return fmt.Errorf("jsonMarshallError: %s", err)
	}

	err = ioutil.WriteFile(filepath.Join(configPath, configFileName+"."+configFileType), configJSON, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	return nil
}

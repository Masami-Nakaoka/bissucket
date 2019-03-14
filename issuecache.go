package main

import (
	"io/ioutil"

	"github.com/namahu/bissucket/config"
)

func readCache() ([]byte, error) {
	cachePath := config.GetConfigValueByKey("issueCachePath")

	return ioutil.ReadFile(cachePath)
}

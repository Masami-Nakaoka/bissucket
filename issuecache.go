package main

import (
	"io/ioutil"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
)

func readCache() ([]byte, error) {
	cachePath := config.GetConfigValueByKey("issueCachePath")

	return ioutil.ReadFile(cachePath)
}

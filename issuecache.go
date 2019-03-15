package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/namahu/bissucket/config"
	bitbucket "github.com/namahu/bissucket/lib"
)

func readCache() ([]byte, error) {
	cachePath := config.GetConfigValueByKey("issueCachePath")
	return ioutil.ReadFile(cachePath)
}

func loadCache(issueCache *IssueCache) error {
	buf, err := readCache()
	if err != nil {
		userName := config.GetConfigValueByKey("bitbucketUserName")
		endPoint := "repositories/" + userName + "/" + issueCache.Repository + "/issues"

		res, err := bitbucket.DoGet(endPoint, userName)
		if err != nil {
			return fmt.Errorf("FetchError: %s", err)
		}

		defer res.Body.Close()

		if err = json.NewDecoder(res.Body).Decode(&issueCache.Store); err != nil {
			return fmt.Errorf("ErrorFF: %s", err)
		}
		return nil
	}

	if err := json.Unmarshal(buf, &issueCache); err != nil {
		return fmt.Errorf("UnmarshalError: %s", err)
	}

	return nil
}

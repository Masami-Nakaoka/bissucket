package issue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"bitbucket.org/Masami_Nakaoka/bissucket/config"
	bitbucket "bitbucket.org/Masami_Nakaoka/bissucket/lib"
	"github.com/urfave/cli"
)

var issueCacheFilePath = os.Getenv("HOME") + "/bissucket.issuecache.json"

func saveIssuesInCache(filePath string, issue *Issues) error {

	buf, err := json.MarshalIndent(issue, "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshallError: %s", err)
	}

	err = ioutil.WriteFile(filePath, buf, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileerror: %s", err)
	}

	return nil
}

func Sync(c *cli.Context) error {

	userName := c.App.Metadata["bitbucketUserName"].(string)
	repositoryName := config.GetConfigValueByKey("defaultRepository")

	endPoint := "repositories/" + userName + "/" + repositoryName + "/issues"

	fmt.Println("Start Issue sync.")

	res, err := bitbucket.DoGet(endPoint, userName)
	if err != nil {
		return fmt.Errorf("APIError: %s", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		return fmt.Errorf("JsonDecodeError: %s", err)
	}

	err = saveIssuesInCache(issueCacheFilePath, issues)
	if err != nil {
		return err
	}

	fmt.Println("Issue synchronization succeeded.")

	return nil
}

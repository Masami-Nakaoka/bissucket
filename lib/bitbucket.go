package bitbucket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/namahu/bissucket/config"
)

const (
	bitbucketURI = "https://api.bitbucket.org/2.0/"
)

func DoGet(endPoint string, userName string) (*http.Response, error) {

	endPoint = bitbucketURI + endPoint

	pass := config.GetConfigValueByKey("bitbucketPassword")

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return nil, fmt.Errorf("RequestError: %s", err)
	}

	req.Header.Set("Contents-type", "application/json")
	req.SetBasicAuth(userName, pass)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ResponseError: %s", err)
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	return res, nil
}

func DoPost(endPoint string, userName string, params map[string]interface{}) error {

	endPoint = bitbucketURI + endPoint
	pass := config.GetConfigValueByKey("bitbucketPassword")

	body, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("RequestError: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(userName, pass)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("ResponseError: %s", err)
	}

	if res.StatusCode != 201 {
		return errors.New(res.Status)
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&res)
}

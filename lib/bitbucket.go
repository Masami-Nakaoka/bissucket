package bitbucket

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/namahu/bissucket/config"
)

const (
	bitbucketURI = "https://api.bitbucket.org/2.0/"
)

var pass = config.GetConfigValueByKey("bitbucketPassword")

func DoGet(endPoint string, userName string) (*http.Response, error) {

	endPoint = bitbucketURI + endPoint

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

func DoPost(endPoint string, userName string, params interface{}) (*http.Response, error) {

	endPoint = bitbucketURI + endPoint
	body := bytes.NewBufferString(params.(string))
	req, err := http.NewRequest("POST", endPoint, body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

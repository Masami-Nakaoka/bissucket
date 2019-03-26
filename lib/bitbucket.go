package bitbucket

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

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

func DoPost(endPoint string, userName string, params url.Values) (*http.Response, error) {

	endPoint = bitbucketURI + endPoint
	pass := config.GetConfigValueByKey("bitbucketPassword")

	var body io.Reader
	body = strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", endPoint, body)
	if err != nil {
		return nil, fmt.Errorf("RequestError: %s", err)
	}
	req.Header.Set("Content-Type", "application/application/x-www-form-urlencoded")
	req.SetBasicAuth(userName, pass)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ResponseError: %s", err)
	}

	if res.StatusCode != 201 {
		return nil, errors.New(res.Status)
	}

	return res, nil
}

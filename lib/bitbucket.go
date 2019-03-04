package bitbucket

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	bitbucketURI = "https://api.bitbucket.org/2.0/"
)

func DoGet(endPoint string, userName string, pass string) (*http.Response, error) {

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

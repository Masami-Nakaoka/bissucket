package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

type Repo struct {
	Size   int `json:"size"`
	Page   int `json:"page"`
	Values []struct {
		// 		Description string `json:"description"`
		// 		IsPrivate   bool   `json:"is_private"`
		// 		Slug        string `json:"slug"`
		// 		Language    string `json:"language"`
		// 		UUID        string `json:"uuid"`
		// 		ForkPolicy  string `json:"fork_policy"`
		// 		Links       struct {
		// 			Pullrequests struct {
		// 				Href string `json:"href"`
		// 			} `json:"pullrequests"`
		// 			Downloads struct {
		// 				Href string `json:"href"`
		// 			} `json:"downloads"`
		// 			Forks struct {
		// 				Href string `json:"href"`
		// 			} `json:"forks"`
		// 			Hooks struct {
		// 				Href string `json:"href"`
		// 			} `json:"hooks"`
		// 			Avatar struct {
		// 				Href string `json:"href"`
		// 			} `json:"avatar"`
		// 			Watchers struct {
		// 				Href string `json:"href"`
		// 			} `json:"watchers"`
		// 			Branches struct {
		// 				Href string `json:"href"`
		// 			} `json:"branches"`
		// 			Tags struct {
		// 				Href string `json:"href"`
		// 			} `json:"tags"`
		// 			Commits struct {
		// 				Href string `json:"href"`
		// 			} `json:"commits"`
		// 			Clone []struct {
		// 				Name string `json:"name"`
		// 				Href string `json:"href"`
		// 			} `json:"clone"`
		// 			Self struct {
		// 				Href string `json:"href"`
		// 			} `json:"self"`
		// 			Source struct {
		// 				Href string `json:"href"`
		// 			} `json:"source"`
		// 			HTML struct {
		// 				Href string `json:"href"`
		// 			} `json:"html"`
		// 		} `json:"links"`
		Name string `json:"name"`
		// 		HasWiki bool   `json:"has_wiki"`
		// 		Website string `json:"website"`
		// 		Scm     string `json:"scm"`
		// 		// 		CreatedOn  time.Time `json:"created_on"`
		// 		Mainbranch struct {
		// 			Name string `json:"name"`
		// 			Type string `json:"type"`
		// 		} `json:"mainbranch"`
		// 		FullName  string `json:"full_name"`
		// 		HasIssues bool   `json:"has_issues"`
		// 		Owner     struct {
		// 			UUID     string `json:"uuid"`
		// 			Type     string `json:"type"`
		// 			Nickname string `json:"nickname"`
		// 			Links    struct {
		// 				Avatar struct {
		// 					Href string `json:"href"`
		// 				} `json:"avatar"`
		// 				HTML struct {
		// 					Href string `json:"href"`
		// 				} `json:"html"`
		// 				Self struct {
		// 					Href string `json:"href"`
		// 				} `json:"self"`
		// 			} `json:"links"`
		// 			AccountID   string `json:"account_id"`
		// 			DisplayName string `json:"display_name"`
		// 			Username    string `json:"username"`
		// 		} `json:"owner"`
		// 		UpdatedOn time.Time `json:"updated_on"`
		// 		Size int    `json:"size"`
		// 		Type string `json:"type"`
	} `json:"values"`
	Pagelen int    `json:"pagelen"`
	Next    string `json:"next"`
}

const (
	baseURL = "https://api.bitbucket.org/2.0/repositories/"
)

var repositories *Repo

func Repository(c *cli.Context) error {
	if c.NArg() > 1 {
		return errors.New("Too many arguments")
	}

	if c.Bool("s") {
		fmt.Print("Sync start")

		repositories, err := getRepositories(c)
		if err != nil {
			return fmt.Errorf("getRepositoriesError: %s", err)
		}

		saveRepository(repositoryCachePath, repositories)

	}

	return nil
}

func getRepositories(c *cli.Context) (*Repo, error) {
	userName := c.App.Metadata["bitbucketUserName"].(string)
	pass := c.App.Metadata["bitbucketPassword"].(string)

	endPoint := bitbucketURI + "repositories/" + userName + "?pagelen=100"

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return nil, fmt.Errorf("RequestError: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(userName, pass)

	fmt.Println("Request Set")
	fmt.Println(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ResponceError: %s", err)
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&repositories)
	if err != nil {
		return nil, fmt.Errorf("JsonDecodeError: %s", err)
	}

	return repositories, nil

}

func saveRepository(filename string, r *Repo) {
	fmt.Println(r)
	buf, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		fmt.Printf("JsonMarshallError: %s", err)

	}

	err = ioutil.WriteFile(filename, buf, os.ModePerm)
	if err != nil {
		fmt.Printf("WriteFileError: %s", err)

	}

}

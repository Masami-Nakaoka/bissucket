package main

import (
	"fmt"

	"github.com/urfave/cli"
)

type Repository struct {
	Size   int `json:"size"`
	Page   int `json:"page"`
	Values []struct {
		Description string `json:"description"`
		IsPrivate   bool   `json:"is_private"`
		Slug        string `json:"slug"`
		Language    string `json:"language"`
		UUID        string `json:"uuid"`
		ForkPolicy  string `json:"fork_policy"`
		Links       struct {
			Pullrequests struct {
				Href string `json:"href"`
			} `json:"pullrequests"`
			Downloads struct {
				Href string `json:"href"`
			} `json:"downloads"`
			Forks struct {
				Href string `json:"href"`
			} `json:"forks"`
			Hooks struct {
				Href string `json:"href"`
			} `json:"hooks"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			Watchers struct {
				Href string `json:"href"`
			} `json:"watchers"`
			Branches struct {
				Href string `json:"href"`
			} `json:"branches"`
			Tags struct {
				Href string `json:"href"`
			} `json:"tags"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Clone []struct {
				Name string `json:"name"`
				Href string `json:"href"`
			} `json:"clone"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Source struct {
				Href string `json:"href"`
			} `json:"source"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
		Name       string    `json:"name"`
		HasWiki    bool      `json:"has_wiki"`
		Website    string    `json:"website"`
		Scm        string    `json:"scm"`
		CreatedOn  time.Time `json:"created_on"`
		Mainbranch struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"mainbranch"`
		FullName  string `json:"full_name"`
		HasIssues bool   `json:"has_issues"`
		Owner     struct {
			UUID     string `json:"uuid"`
			Type     string `json:"type"`
			Nickname string `json:"nickname"`
			Links    struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"links"`
			AccountID   string `json:"account_id"`
			DisplayName string `json:"display_name"`
			Username    string `json:"username"`
		} `json:"owner"`
		UpdatedOn time.Time `json:"updated_on"`
		Size      int       `json:"size"`
		Type      string    `json:"type"`
	} `json:"values"`
	Pagelen int `json:"pagelen"`
}

const (
	baseUrl = "https//api.bitbucket.org/2.0/repositories/"
)

func Repository(c *cli.Context) error {
	userName := c.App.Metadata["bitbucketUserName"]
	token := c.App.Metadata["bitbucketToken"]

	endPoint := baseUrl + userName + "/"

	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return err
	}
	return nil
}

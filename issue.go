package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	bitbucket "bitbucket.org/bissucket/lib"

	"github.com/urfave/cli"
)

type Issues struct {
	Pagelen int `json:"pagelen"`
	Page    int `json:"page"`
	Size    int `json:"size"`
	Values  []struct {
		Priority   string `json:"priority"`
		Kind       string `json:"kind"`
		Repository struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Type     string `json:"type"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			UUID     string `json:"uuid"`
		} `json:"repository"`
		Links struct {
			Attachments struct {
				Href string `json:"href"`
			} `json:"attachments"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Watch struct {
				Href string `json:"href"`
			} `json:"watch"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Vote struct {
				Href string `json:"href"`
			} `json:"vote"`
		} `json:"links"`
		Reporter struct {
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			AccountID   string `json:"account_id"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Nickname string `json:"nickname"`
			Type     string `json:"type"`
			UUID     string `json:"uuid"`
		} `json:"reporter"`
		Title     string      `json:"title"`
		Component interface{} `json:"component"`
		Votes     int         `json:"votes"`
		Watches   int         `json:"watches"`
		Content   struct {
			Raw    string `json:"raw"`
			Markup string `json:"markup"`
			HTML   string `json:"html"`
			Type   string `json:"type"`
		} `json:"content"`
		Assignee struct {
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			AccountID   string `json:"account_id"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Nickname string `json:"nickname"`
			Type     string `json:"type"`
			UUID     string `json:"uuid"`
		} `json:"assignee"`
		State     string      `json:"state"`
		Version   interface{} `json:"version"`
		EditedOn  interface{} `json:"edited_on"`
		CreatedOn time.Time   `json:"created_on"`
		Milestone interface{} `json:"milestone"`
		UpdatedOn time.Time   `json:"updated_on"`
		Type      string      `json:"type"`
		ID        int         `json:"id"`
	}
}

var issues *Issues

func Issue(c *cli.Context) error {
	if c.NArg() > 1 {
		return errors.New("Too manu arguments.")
	} else if c.NArg() == 0 {

	}

	userName := c.App.Metadata["bitbucketUserName"].(string)
	pass := c.App.Metadata["bitbucketPassword"].(string)

	repositoryName := c.Args().First()

	res, err := fecthAllIssueFromBitbucket(repositoryName, userName, pass)
	if err != nil {
		return fmt.Errorf("FetchError: %s", err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		return fmt.Errorf("DecodeError: %s", err)
	}

	if c.Int("d") > 0 {
		issueID := c.Int("d")

		fmt.Println("------------------------------")
		fmt.Println("Details of " + repositoryName + "'s issue")
		fmt.Println("------------------------------")

		for _, issue := range issues.Values {
			if issueID == issue.ID {
				fmt.Print("Issue ID: ")
				fmt.Println(issue.ID)
				fmt.Println("Issue Title: " + issue.Title)
				fmt.Println("Issue Description: ")
				fmt.Println(issue.Content.Raw)
				fmt.Println("")
			}
		}

	} else {
		fmt.Println("------------------------------")
		fmt.Println("Issue List of " + repositoryName)
		fmt.Println("------------------------------")
		fmt.Println("ID / Title / Type / State / Priority / Kind / Assignee")

		var issueTemplate string
		for _, issue := range issues.Values {
			issueTemplate = strconv.Itoa(issue.ID) + " / " + issue.Title + " / " + issue.Type + " / " + issue.State + " / " + issue.Priority + " / " + issue.Kind + " / " + issue.Assignee.Username
			fmt.Println(issueTemplate)
		}
	}

	return nil
}

func fecthAllIssueFromBitbucket(repositoryName string, userName string, pass string) (*http.Response, error) {
	endPoint := bitbucketURI + "repositories/" + userName + "/" + repositoryName + "/issues"

	res, err := bitbucket.DoGet(endPoint, userName, pass)
	if err != nil {
		return nil, fmt.Errorf("APIError: %s", err)
	}

	return res, nil

}

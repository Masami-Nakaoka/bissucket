package main

import (
	"fmt"
	"time"

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

func Issue(c *cli.Context) error {
	if c.Bool("l") {
		err := showIssuesList()
		if err != nil {
			return fmt.Errorf("showIssuesListErrot: %s", err)
		}
	}

	return nil
}

func showIssuesList() error {
	return nil
}

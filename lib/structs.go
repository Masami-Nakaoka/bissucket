package bitbucket

import "time"

type Repos struct {
	// Size   int `json:"size"`
	// Page   int `json:"page"`
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
		HasIssues bool `json:"has_issues"`
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
	// Pagelen int    `json:"pagelen"`
	// Next    string `json:"next"`
}

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

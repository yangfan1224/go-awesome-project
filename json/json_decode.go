package json

import (
	"time"
	"net/url"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	"os"
	"text/template"
)

const IssuesURL = "https://api.github.com/search/issues"

const templ = `{{.TotalCount}} issues:
{{range.Items}}------------------------------------------
	Url: 				{{.Url}}				
	RepositoryUrl:  	{{.RepositoryUrl}}	
	LabelsUrl: 			{{.LabelsUrl}} 
	CommentsUrl:		{{.CommentsUrl}}
	EventsUrl:			{{.EventsUrl}} 		
	HTMLUrl: 			{{.HTMLUrl}}
	ID: 				{{.ID}}
	NodeID: 			{{.NodeID}}
	Number: 	    	{{.Number}}
	Title: 				{{.Title}}
	User: 		    	{{.User.Login}}  			
	Labels: 			{{.Labels}}
	State: 				{{.State}}
	Locked: 			{{.Locked}}
	Assignee: 			{{.Assignee}}
	Assignees: 			{{.Assignees}}
	Milestone: 			{{.Milestone}}
	Comments: 			{{.Comments}}
	CreatedAt: 			{{.CreatedAt}}
	UpdatedAt: 			{{.UpdatedAt}}
	ClosedAt: 			{{.ClosedAt}}
	AuthorAssociation: 	{{.AuthorAssociation}}
	Body: 				{{.Body}}
	Score: 				{{.Score}}
{{end}}
`
type User struct {
	Login 					string 	`json:"login"`
	ID 						int 	`json:"id"`
	NodeID 					string `json:"node_id"`
	AvatarUrl 				string `json:"avatar_url"`
	GravatarID 				string `json:"gravatar_id"`
	url 					string `json:"url"`
	HTMLUrl 				string `json:"html_url"`
	FollowersUrl 			string `json:"followers_url"`
	FollowingUrl 			string `json:"following_url"`
	GistsUrl 				string `json:"gists_url"`
	StarredUrl 				string `json:"starred_url"`
	SubscriptionsUrl 		string `json:"subscriptions_url"`
	OrganizationsUrl 		string `json:"organizations_url"`
	ReposUrl 				string `json:"repos_url"`
	EventsUrl 				string `json:"events_url"`
	ReceivedEventsUrl 		string `json:"received_events_url"`
	Type 					string `json:"type"`
	SiteAdmin 				bool 	`json:"site_admin"`
}

type Assignee struct {
	Login 					string `json:"login"`
	ID 						int    `json:"id"`
	NodeID 					string `json:"node_id"`
	AvatarUrl 				string `json:"avatar_url"`
	GravatarID 				string `json:"gravatar_id"`
	url 					string `json:"url"`
	HTMLUrl 				string `json:"html_url"`
	FollowersUrl 			string `json:"followers_url"`
	FollowingUrl 			string `json:"following_url"`
	GistsUrl 				string `json:"gists_url"`
	StarredUrl 				string `json:"starred_url"`
	SubscriptionsUrl 		string `json:"subscriptions_url"`
	OrganizationsUrl 		string `json:"organizations_url"`
	ReposUrl 				string `json:"repos_url"`
	EventsUrl 				string `json:"events_url"`
	ReceivedEventsUrl 		string `json:"received_events_url"`
	Type 					string `json:"type"`
	SiteAdmin 				bool   `json:"site_admin"`
}

type Assignees struct {
	Login 					string `json:"login"`
	ID 						int    `json:"id"`
	NodeID 					string `json:"node_id"`
	AvatarUrl 				string `json:"avatar_url"`
	GravatarID 				string `json:"gravatar_id"`
	url 					string `json:"url"`
	HTMLUrl 				string `json:"html_url"`
	FollowersUrl 			string `json:"followers_url"`
	FollowingUrl 			string `json:"following_url"`
	GistsUrl 				string `json:"gists_url"`
	StarredUrl 				string `json:"starred_url"`
	SubscriptionsUrl 		string `json:"subscriptions_url"`
	OrganizationsUrl 		string `json:"organizations_url"`
	ReposUrl 				string `json:"repos_url"`
	EventsUrl 				string `json:"events_url"`
	ReceivedEventsUrl 		string `json:"received_events_url"`
	Type 					string `json:"type"`
	SiteAdmin 				bool   `json:"site_admin"`
}
type Label struct {
	ID 			int 	`json:"id"`
	NodeID 		string `json:"node_id"`
	Url 		string `json:"url"`
	Name 		string `json:"name"`
	Color 		string `json:"color"`
	Default 	bool 	`json:"default"`
}
type Item struct {
	Url 				string 			`json:"url"`
	RepositoryUrl 		string 			`json:"repository_url"`
	LabelsUrl 			string 			`json:"labels_url"`
	CommentsUrl 		string 			`json:"comments_url"`
	EventsUrl 			string 			`json:"events_url"`
	HTMLUrl 			string 			`json:"html_url"`
	ID 					int    			`json:"id"`
	NodeID 				string 			`json:"node_id"`
	Number 				int    			`json:"number"`
	Title 				string 			`json:"title"`
	User 				*User  			`json:"user"`
	Labels 				[]*Label 		`json:"labels"`
	State 				string 			`json:"state"`
	Locked 				bool   			`json:"locked"`
	Assignee 			*Assignee 		`json:"assignee"`
	Assignees 			[]*Assignees	`json:"assignees"`
	Milestone 			*string 		`json:"milestone"`
	Comments 			int 			`json:"comments"`
	CreatedAt 			time.Time 		`json:"created_at"`
	UpdatedAt 			time.Time 		`json:"updated_at"`
	ClosedAt 			*string 		`josn:"closed_at"`
	AuthorAssociation 	string 			`josn:"author_association"`
	Body 				string 			`json:"body"`
	Score 				float32 		`json:"score"`
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count""`
	InCompleteResults bool `json:"incomplete_results"`
	Items []*Item `json:items`
}

func (issues *IssuesSearchResult) String() (string){
	var items string
	for _, item := range issues.Items{
		items += fmt.Sprintf("itme=%v", item)
	}
	return fmt.Sprintf("TotalCount=%d, InCompleteResults=%t, Items = %s", issues.TotalCount, issues.InCompleteResults, items)
}

func SearchIssues(terms []string) (*IssuesSearchResult, error){
	q := url.QueryEscape(strings.Join(terms," "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		fmt.Errorf("http get %s failed", IssuesURL + "?q=" + q)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	//json.NewEncoder(os.Stdout).Encode(result)
	TemplateReport(&result)
	return &result, nil
}


func TemplateReport(issues *IssuesSearchResult)(error){
	report, err := template.New("issuelist").Parse(templ)
	if err != nil{
		fmt.Errorf("TemplateReport failed: %s", err)
		return err
	}

	report.Execute(os.Stdout, issues)
	return nil
}

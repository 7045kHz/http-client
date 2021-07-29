package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/7045kHz/http-client/gohttp"
)

type GitHub struct {
	CurrentUserURL                   string `json:"current_user_url"`
	CurrentUserAuthorizationsHTMLURL string `json:"current_user_authorizations_html_url"`
	AuthorizationsURL                string `json:"authorizations_url"`
	CodeSearchURL                    string `json:"code_search_url"`
	CommitSearchURL                  string `json:"commit_search_url"`
	EmailsURL                        string `json:"emails_url"`
	EmojisURL                        string `json:"emojis_url"`
	EventsURL                        string `json:"events_url"`
	FeedsURL                         string `json:"feeds_url"`
	FollowersURL                     string `json:"followers_url"`
	FollowingURL                     string `json:"following_url"`
	GistsURL                         string `json:"gists_url"`
	HubURL                           string `json:"hub_url"`
	IssueSearchURL                   string `json:"issue_search_url"`
	IssuesURL                        string `json:"issues_url"`
	KeysURL                          string `json:"keys_url"`
	LabelSearchURL                   string `json:"label_search_url"`
	NotificationsURL                 string `json:"notifications_url"`
	OrganizationURL                  string `json:"organization_url"`
	OrganizationRepositoriesURL      string `json:"organization_repositories_url"`
	OrganizationTeamsURL             string `json:"organization_teams_url"`
	PublicGistsURL                   string `json:"public_gists_url"`
	RateLimitURL                     string `json:"rate_limit_url"`
	RepositoryURL                    string `json:"repository_url"`
	RepositorySearchURL              string `json:"repository_search_url"`
	CurrentUserRepositoriesURL       string `json:"current_user_repositories_url"`
	StarredURL                       string `json:"starred_url"`
	StarredGistsURL                  string `json:"starred_gists_url"`
	UserURL                          string `json:"user_url"`
	UserOrganizationsURL             string `json:"user_organizations_url"`
	UserRepositoriesURL              string `json:"user_repositories_url"`
	UserSearchURL                    string `json:"user_search_url"`
}
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	githubClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()
	client.SetConnectionTimeout(2 * time.Second)
	client.SetRequestTimeout(50 * time.Millisecond)
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json ; charset=utf-8")
	client.SetHeaders(commonHeaders)
	return client
}
func main() {
	getGitAPI()
	//createUser()

}
func createUser(user User) {

	response, err := githubClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read: ", err)
	}
	var p GitHub
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}
	j, _ := json.MarshalIndent(p, "", "    ")
	fmt.Printf("%v\n", response.StatusCode)
	fmt.Printf("%s\n\n\n\n", j)
	fmt.Printf("%s\n", p.UserURL)

}

func getGitAPI() {

	response, err := getGithubClient().Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read: ", err)
	}
	var p GitHub
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}
	j, _ := json.MarshalIndent(p, "", "    ")
	fmt.Printf("%v\n", response.StatusCode)
	fmt.Printf("%s\n\n\n\n", j)
	fmt.Printf("%s\n", p.UserURL)

}
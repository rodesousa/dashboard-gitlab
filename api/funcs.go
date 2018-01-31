package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repository struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type MergeRequest struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	WebUrl string `json:"web_url"`
}

type RepositoryState struct {
	MergeRequests []MergeRequest
	Repository
}

var (
	Token           string = ""
	MergeRequests   map[int]MergeRequest
	RepositoriesMap map[int]*RepositoryState
)

func call_api(url string) []byte {
	get_url, _ := http.Get(url)
	defer get_url.Body.Close()
	body, _ := ioutil.ReadAll(get_url.Body)
	return body
}

func add_merge_request(repo string, mr MergeRequest) {
	MergeRequests[mr.Id] = mr
}

func get_merge_request(id int) []MergeRequest {
	body := call_api(
		fmt.Sprintf(
			"https://gitlab.com/api/v4/projects/%d/merge_requests?private_token=%s&state=opened",
			id, Token))
	var mrs []MergeRequest
	json.Unmarshal(body, &mrs)
	return mrs
}

func Init_MergeRequestsInfo() {
	MergeRequests = make(map[int]MergeRequest)
}

func Maj_state() {
	dat, _ := ioutil.ReadFile("/tmp/test.json")
	var rep []Repository
	json.Unmarshal(dat, &rep)
	RepositoriesMap = make(map[int]*RepositoryState)
	for _, element := range rep {
		RepositoriesMap[element.Id] = &RepositoryState{make([]MergeRequest, 0), Repository{element.Name, element.Id}}
	}
}

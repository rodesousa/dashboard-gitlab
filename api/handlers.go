package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Add_repo(id string) {
	//todo
	return
}

func All_merge_requests() gin.H {
	var repos = make([]RepositoryState, 0, len(RepositoriesMap))
	fmt.Println(RepositoriesMap)
	for _, element := range RepositoriesMap {
		if len(element.MergeRequests) > 0 {
			repos = append(repos, *element)
		}
	}
	return gin.H{
		"merge_requests": repos,
	}
}

func Maj_merge_request() (int, gin.H) {
	for _, mrs := range RepositoriesMap {
		tmp := get_merge_request(mrs.Id)
		mrs.MergeRequests = tmp
	}
	return 200, gin.H{"status": "ok"}
}

func Save_state() (int, gin.H) {
	rep := make([]Repository, 0, len(RepositoriesMap))
	for _, element := range RepositoriesMap {
		rep = append(rep, Repository{element.Name, element.Id})
	}
	repo, _ := json.Marshal(rep)
	ioutil.WriteFile("/tmp/test.json", repo, 0644)
	return 200, gin.H{"status": "ok"}
}

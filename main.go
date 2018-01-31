package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdesousa/dashboard-gitlab/api"
)

func main() {
	r := gin.Default()

	// init
	api.Init_MergeRequestsInfo()
	api.Maj_state()
	go api.Maj_merge_request()
	r.StaticFS("/static", http.Dir("frontend/build/static"))
	r.LoadHTMLFiles("frontend/build/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/merge_request/all", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, api.All_merge_requests())
	})

	r.GET("/repo/save", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(api.Save_state())
	})

	r.GET("/merge_request/maj", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(api.Maj_merge_request())
	})

	r.POST("/repo/add", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})

	r.POST("/token", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		value := c.Query("value")
		api.Token = value
	})
	// r.GET("/api/merge_request", api.Merge_Request)

	r.Run() // listen and serve on 0.0.0.0:8080

}

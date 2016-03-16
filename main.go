package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	root := gin.Default()
	root.LoadHTMLGlob("src/templates/*")

	root.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "FOOBAR",
		})
	})

	root.Static("assets", "dist")

	v1 := root.Group("/v1")
	{
		v1.POST("/login", nil)
	}

	err := root.RunTLS(":4499", "conf/cert.pem", "conf/server.key")
	if err != nil {
		fmt.Println(err)
	}
}

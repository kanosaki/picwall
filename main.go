package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanosaki/picwall/server"
	"net/http"
	"os"
)

func main() {
	app := server.NewApp()
	devImagesDir := os.ExpandEnv("$IMAGES_DIR")
	root := gin.Default()
	root.LoadHTMLGlob("src/templates/*")

	root.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "FOOBAR",
		})
	})

	root.Static("assets", "dist")
	root.Static("icons", "node_modules/material-design-icons")

	// dev endpoints
	devEp := root.Group("/dev")
	{
		devEp.GET("/ws", server.SinkHandler(app))
		devEp.Static("image", devImagesDir)
	}

	v1 := root.Group("/v1")
	{
		v1.POST("/login", nil)
		// tap := v1.Group("/tap")
	}

	err := root.RunTLS(":4499", "conf/cert.pem", "conf/server.key")
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/cloudfoundry-incubator/candiedyaml"
	"github.com/gin-gonic/gin"
	"github.com/kanosaki/picwall/server"
	"github.com/kanosaki/picwall/server/core"
	"github.com/kanosaki/picwall/server/source"
	"net/http"
	"os"
)

func readconfig() (*core.Config, error) {
	f, err := os.Open("conf/default.yaml")
	if err != nil {
		return nil, err
	}
	decoder := candiedyaml.NewDecoder(f)
	var ret core.Config
	err = decoder.Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func initSource(app *core.App, config map[string]core.SourceConfig) (*source.BottomSource, error) {
	sources := make([]core.Source, 0, len(config))
	tumblrTr := core.NewTumblrTransformer(app.Tumblr)
	twitterTr := core.NewTweetTransformer(app.Twitter)
	for key, cfg := range config {
		logrus.Infof("Loading source %s(%s)...", key, cfg.Type())
		switch cfg.Type() {
		case "twitter/list":
			listname := cfg.Get("listname").(string)
			twitterSrc, err := source.NewTwitterListSource(app.Twitter, listname, twitterTr)
			if err != nil {
				return nil, err
			}
			sources = append(sources, twitterSrc)
		case "tumblr/dashboard":
			tumblrSrc, err := source.NewTumblrDashboard(app.Tumblr, tumblrTr)
			if err != nil {
				return nil, err
			}
			sources = append(sources, tumblrSrc)
		}
	}
	return source.NewBottomSource(sources, 100), nil
}

func main() {
	config, err := readconfig()
	if err != nil {
		logrus.Fatalf("Config not loaded: %s", err)
	}
	app, err := core.NewApp(config)
	if err != nil {
		logrus.Fatalf("Initialization failed: %s", err)
	}
	src, err := initSource(app, config.Sources)
	if err != nil {
		logrus.Fatalf("Source initialization failed: %s", err)
	}
	for entry := range src.Faucet() {
		fmt.Printf("%s %s\n", entry.Get("id", "==="), entry.String())
	}
	return
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

	err = root.RunTLS(":4499", "conf/cert.pem", "conf/server.key")
	if err != nil {
		fmt.Println(err)
	}
}

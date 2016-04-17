package server

import (
	"github.com/kanosaki/picwall/model"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	"github.com/k0kubun/pp"
)

type SinkRequest struct {
	Count int `json:"count"`
	ReqId int `json:"reqId"`
}

type SinkResponse struct {
	ReqId    int `json:"reqId"`
	Contents []model.PicItem `json:"contents"`
}

func SinkHandler(wsUpgrader websocket.Upgrader) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
		if err != nil {
			c.Error(err)
			return
		}
		devImagesDir := os.ExpandEnv("$IMAGES_DIR")
		files, err := ioutil.ReadDir(devImagesDir)
		if err != nil {
			c.Error(err)
		}
		retItems := make([]model.PicItem, 0, len(files))
		for _, f := range files {
			fName := f.Name()
			if strings.HasSuffix(fName, ".png") || strings.HasSuffix(fName, ".jpg") {
				retItems = append(retItems, model.PicItem{
					Caption: fName,
					ThumbnailURL: fmt.Sprintf("/dev/image/%s", fName),
				})
			}
		}
		cursor := 0
		for {
			var req SinkRequest
			err = conn.ReadJSON(&req)
			if err != nil {
				pp.Println(err)
				break
			}
			endCursor := cursor + req.Count
			if endCursor >= len(retItems) {
				endCursor = len(retItems) - 1
			}
			ret := SinkResponse{
				ReqId: req.ReqId,
				Contents: retItems[cursor:endCursor],
			}
			err = conn.WriteJSON(ret)
			cursor = endCursor
			if err != nil {
				pp.Println(err)
				break
			}
		}
	}
}

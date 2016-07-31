package source

import (
	"container/list"
	"io/ioutil"
	"fmt"
	"github.com/kanosaki/picwall/server/core"
)

type DirectorySrouce struct {
	path   string
	ch     core.Stream
	buffer *list.List // FileInfo
}

func NewDirectorySource(path string) (*DirectorySrouce, error) {
	ds := &DirectorySrouce{
		path: path,
		ch: make(core.Stream),
		buffer: list.New(),
	}
	files, err := ioutil.ReadDir(path)
	buffer := list.New()
	if err != nil {
		return nil, err
	} else if len(files) == 0 {
		return nil, fmt.Errorf("Emptyr directory is not allowed as direcotry source: %s", path)
	}
	for _, f := range files {
		buffer.PushBack(f)
	}
	return ds, nil
}

func (ds *DirectorySrouce) Drain(count int) {
	doneCount := 0
	for doneCount < count || ds.buffer.Len() != 0 {
		fInfo := ds.buffer.Remove(ds.buffer.Front())
		doneCount += 1
	}
}

package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Entry struct {
	data map[string]interface{}
}

func NewEntry() *Entry {
	return &Entry{
		data: make(map[string]interface{}),
	}
}

func (e *Entry) WriteJons(w http.ResponseWriter) error {
	enc := json.NewEncoder(w)
	return enc.Encode(e.data)
}

func (e *Entry) Get(key string, defaultValue interface{}) interface{} {
	if v, ok := e.data[key]; ok {
		return v
	} else {
		return defaultValue
	}
}

func (e *Entry) Set(key string, value interface{}) {
	e.data[key] = value
}

func (e *Entry) Has(key string) bool {
	_, ok := e.data[key]
	return ok
}

func (e *Entry) String() string {
	return fmt.Sprintf("%v", e.data)
}

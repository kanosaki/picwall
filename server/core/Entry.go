package core

import (
	"net/http"
	"encoding/json"
)

type Entry struct {
	data map[string]interface{}
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


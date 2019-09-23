package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type answer struct {
	w    http.ResponseWriter
	enc  *json.Encoder
	init sync.Once
}

func (resp *answer) json(obj interface{}, code int) error {
	if resp.enc == nil {
		resp.init.Do(func() {
			resp.enc = json.NewEncoder(resp.w)
		})
	}
	resp.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.w.WriteHeader(code)
	return resp.enc.Encode(obj)
}

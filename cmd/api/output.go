package main

import (
	"encoding/json"
	"net/http"
)

type ctxkey string

const outputKey = ctxkey("output")

type output struct {
	Payload interface{} `json:"payload"`
	Code    int         `json:"code"`
}

// it is not a middleware - signature is intended
func (outval *output) into(r *http.Request, w http.ResponseWriter) {

	out, ok := r.Context().Value(outputKey).(*output)
	if out == nil || !ok {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("poor context")
		return
	}

	if outval.empty() {
		outval.Code = http.StatusNoContent
	}

	*out = *outval
}

func (outval *output) empty() bool {
	return outval.Code == 0 && outval.Payload == nil
}

func out(msg interface{}, code int) *output {
	return &output{msg, code}
}

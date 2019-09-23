package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func jzon(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		ctx := context.WithValue(
			r.Context(),
			outputKey,
			&output{},
		)
		// this new r is magic
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		enc := json.NewEncoder(w)
		ctx = r.Context()
		out, ok := ctx.Value(outputKey).(*output)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			enc.Encode(output{"no output", http.StatusInternalServerError})
			return
		}
		w.WriteHeader(out.Code)
		enc.Encode(out)
	}
}

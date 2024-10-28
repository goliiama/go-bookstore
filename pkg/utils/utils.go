package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// unmarshalling json data format to usable Go data structure
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

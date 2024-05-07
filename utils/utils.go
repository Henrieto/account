package utils

import (
	"encoding/json"
	"net/http"
)

func BindJson(r *http.Request, object any) (err error) {
	Decoder := json.NewDecoder(r.Body)
	err = Decoder.Decode(object)
	return
}

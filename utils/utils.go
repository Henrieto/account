package utils

import (
	"encoding/json"
	"net/http"
)

// create a custom err
type Error string

func (err Error) Error() string { return string(err) }

var (
	NotString            Error = "not a string"
	NotAMapOfStringToAny Error = "not a map of string to any"
)

func BindJson(r *http.Request, object any) (err error) {
	Decoder := json.NewDecoder(r.Body)
	err = Decoder.Decode(object)
	return
}

func GetString(_object any) (string, error) {
	switch object := _object.(type) {
	case string:
		return object, nil
	default:
		return "", NotString
	}
}

func GetStringList(object any) ([]string, error) {
	switch object := object.(type) {
	case []string:
		return object, nil
	default:
		return nil, NotString
	}
}

func GetMapStringToAny(_object any) (map[string]any, error) {
	switch object := _object.(type) {
	case map[string]any:
		return object, nil
	default:
		return nil, NotAMapOfStringToAny
	}
}

func GetMapStringToString(_object any) (map[string]string, error) {
	switch object := _object.(type) {
	case map[string]string:
		return object, nil
	default:
		return nil, NotAMapOfStringToAny
	}
}

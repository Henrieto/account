package utils

import (
	"net/http"

	db_utils "github.com/henrieto/account/utils/db"
	"github.com/henrieto/jax"
)

type PaginatorResponse struct {
	Data      any                 `json:"data,omitempty"`
	Paginator *db_utils.Paginator `json:"paginator,omitempty"`
}

type Response struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	return new(Response)
}
func (res *Response) Send(writer http.ResponseWriter, status int) error {
	err := jax.Json(writer, res, status)
	if err != nil {
		return err
	}
	return nil
}

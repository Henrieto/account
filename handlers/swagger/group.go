package swagger

import "time"

type GroupData struct {
	Name string `json:"name,omitempty"`
}

type GroupSuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   struct {
		Name     string    `json:"name,omitempty"`
		CreateAt time.Time `json:"create_at,omitempty"`
		UpdateAt time.Time `json:"update_at,omitempty"`
	} `json:"data,omitempty"`
}

type GroupFailed struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any
}

type GroupDeleteSuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any
}

type GroupArraySuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   []struct {
		Name     string    `json:"name,omitempty"`
		CreateAt time.Time `json:"create_at,omitempty"`
		UpdateAt time.Time `json:"update_at,omitempty"`
	} `json:"data,omitempty"`
}

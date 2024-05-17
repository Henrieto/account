package swagger

import (
	"time"
)

type PermissionData struct {
	Model    string `json:"model,omitempty"`
	Name     string `json:"name,omitempty"`
	Codename string `json:"codename,omitempty"`
}

type PermissionsGroupData struct {
	GroupId      int32 `json:"group_id,omitempty"`
	PermissionId int32 `json:"permission_id,omitempty"`
}

type PermissionsUserData struct {
	UserId       int32 `json:"group_id,omitempty"`
	PermissionId int32 `json:"permission_id,omitempty"`
}

type PermissionSuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   struct {
		Model    string    `json:"model,omitempty"`
		Name     string    `json:"name,omitempty"`
		Codename string    `json:"codename,omitempty"`
		CreateAt time.Time `json:"create_at,omitempty"`
		UpdateAt time.Time `json:"update_at,omitempty"`
	} `json:"data,omitempty"`
}

type GetGroupPermissionSuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   struct {
		ID           int32     `json:"id"`
		Model        string    `json:"model"`
		Name         string    `json:"name"`
		Codename     string    `json:"codename"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		ID_2         int32     `json:"id_2"`
		GroupID      int32     `json:"group_id"`
		PermissionID int32     `json:"permission_id"`
		CreatedAt_2  time.Time `json:"created_at_2"`
		UpdatedAt_2  time.Time `json:"updated_at_2"`
	} `json:"data,omitempty"`
}

type GetUserPermissionSuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   struct {
		ID           int32     `json:"id"`
		Model        string    `json:"model"`
		Name         string    `json:"name"`
		Codename     string    `json:"codename"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		ID_2         int32     `json:"id_2"`
		UserID       int32     `json:"group_id"`
		PermissionID int32     `json:"permission_id"`
		CreatedAt_2  time.Time `json:"created_at_2"`
		UpdatedAt_2  time.Time `json:"updated_at_2"`
	} `json:"data,omitempty"`
}

type PermissionArraySuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   []struct {
		Model    string    `json:"model,omitempty"`
		Name     string    `json:"name,omitempty"`
		Codename string    `json:"codename,omitempty"`
		CreateAt time.Time `json:"create_at,omitempty"`
		UpdateAt time.Time `json:"update_at,omitempty"`
	} `json:"data,omitempty"`
}

type PermissionFailed struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type PermissionDeleteSuccess struct {
	Status string `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

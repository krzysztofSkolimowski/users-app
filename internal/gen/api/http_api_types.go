// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// PatchUser defines model for PatchUser.
type PatchUser struct {
	Country   *string              `json:"country,omitempty"`
	Email     *openapi_types.Email `json:"email,omitempty"`
	FirstName *string              `json:"first_name,omitempty"`
	LastName  *string              `json:"last_name,omitempty"`
	Nickname  *string              `json:"nickname,omitempty"`
}

// PostUser defines model for PostUser.
type PostUser struct {
	Country   string              `json:"country"`
	Email     openapi_types.Email `json:"email"`
	FirstName string              `json:"first_name"`
	LastName  string              `json:"last_name"`
	Nickname  string              `json:"nickname"`
	Password  string              `json:"password"`
}

// User defines model for User.
type User struct {
	Country   string              `json:"country"`
	CreatedAt time.Time           `json:"created_at"`
	Email     openapi_types.Email `json:"email"`
	FirstName string              `json:"first_name"`
	Id        openapi_types.UUID  `json:"id"`
	LastName  string              `json:"last_name"`
	Nickname  string              `json:"nickname"`
	UpdatedAt time.Time           `json:"updated_at"`
}

// Users defines model for Users.
type Users struct {
	Users []User `json:"users"`
}

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	FirstName *string              `form:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  *string              `form:"last_name,omitempty" json:"last_name,omitempty"`
	Nickname  *string              `form:"nickname,omitempty" json:"nickname,omitempty"`
	Email     *openapi_types.Email `form:"email,omitempty" json:"email,omitempty"`
	Country   *string              `form:"country,omitempty" json:"country,omitempty"`
	Limit     *int32               `form:"limit,omitempty" json:"limit,omitempty"`
	Offset    *int32               `form:"offset,omitempty" json:"offset,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = PostUser

// PatchUsersUserIDJSONRequestBody defines body for PatchUsersUserID for application/json ContentType.
type PatchUsersUserIDJSONRequestBody = PatchUser

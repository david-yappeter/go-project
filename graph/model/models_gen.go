// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
)

type FileUploadOps struct {
	UploadSingle *FileUpload   `json:"upload_single"`
	UploadBatch  []*FileUpload `json:"upload_batch"`
}

type FileUploadPagination struct {
	Limit     *int          `json:"limit"`
	Page      *int          `json:"page"`
	Asc       *bool         `json:"asc"`
	SortBy    *string       `json:"sort_by"`
	Scopes    *bool         `json:"scopes"`
	UserID    *int          `json:"user_id"`
	TotalData int           `json:"total_data"`
	Nodes     []*FileUpload `json:"nodes"`
}

type IgPostOps struct {
	Create        *IgPost `json:"create"`
	Update        *IgPost `json:"update"`
	Archive       string  `json:"archive"`
	Unarchive     string  `json:"unarchive"`
	SoftDelete    string  `json:"soft_delete"`
	RestoreDelete string  `json:"restore_delete"`
	HardDelete    string  `json:"hard_delete"`
}

type IgPostPagination struct {
	Limit     *int      `json:"limit"`
	Page      *int      `json:"page"`
	Asc       *bool     `json:"asc"`
	SortBy    *string   `json:"sort_by"`
	Scopes    *bool     `json:"scopes"`
	TotalData int       `json:"total_data"`
	Nodes     []*IgPost `json:"nodes"`
}

type NewIgPost struct {
	Caption string            `json:"caption"`
	Files   []*graphql.Upload `json:"files"`
}

type NewIgPostFile struct {
	FileID string `json:"file_id"`
	PostID int    `json:"post_id"`
}

type NewUser struct {
	Name            string  `json:"name"`
	Password        string  `json:"password"`
	Email           string  `json:"email"`
	Address         *string `json:"address"`
	TelephoneNumber *string `json:"telephone_number"`
}

type TokenData struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type TokenOps struct {
	Login           *TokenData `json:"login"`
	RegisterByEmail *TokenData `json:"register_by_email"`
}

type UpdateIgPost struct {
	ID         int     `json:"id"`
	Caption    *string `json:"caption"`
	IsArchived *int    `json:"is_archived"`
}

type UpdateUser struct {
	ID                      int     `json:"id"`
	Name                    string  `json:"name"`
	Password                string  `json:"password"`
	Email                   string  `json:"email"`
	Address                 *string `json:"address"`
	TelephoneNumber         *string `json:"telephone_number"`
	CreatedAt               string  `json:"created_at"`
	UpdatedAt               *string `json:"updated_at"`
	DeletedAt               *string `json:"deleted_at"`
	EmailVerificationHash   *string `json:"email_verification_hash"`
	EmailVerificationStatus int     `json:"email_verification_status"`
}

type UserOps struct {
	Create        *User  `json:"create"`
	Update        *User  `json:"update"`
	SoftDelete    string `json:"soft_delete"`
	RestoreDelete string `json:"restore_delete"`
	Delete        string `json:"delete"`
}

type UserPagination struct {
	Limit     *int    `json:"limit"`
	Page      *int    `json:"page"`
	Asc       *bool   `json:"asc"`
	SortBy    *string `json:"sort_by"`
	Scopes    *bool   `json:"scopes"`
	TotalData int     `json:"total_data"`
	Nodes     []*User `json:"nodes"`
}

type UserRegisterByEmail struct {
	Name            string  `json:"name"`
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
	Address         *string `json:"address"`
	TelephoneNumber *string `json:"telephone_number"`
}

package entity

import "github.com/gofrs/uuid"

type User struct {
	UID uuid.UUID `json:"uuid" dynamo:"uid"`
	Version int `json:"-" dynamo:"document_version"`
	Username string `json:"username" dynamo:"username"`
	Email string `json:"email" dynamo:"email"`
	EmailVerified bool `json:"email_verified" dynamo:"email_verified"`
	PasswordHash []byte `json:"-" dynamo:"password_hash"`
}

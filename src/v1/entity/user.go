package entity

import (
    "github.com/gofrs/uuid"
    "golang.org/x/crypto/bcrypt"
    "time"
)

const (
    UserVersion  = 1
    PasswordCost = 10
)

type User struct {
	UID uuid.UUID `json:"uuid" dynamo:"uid"`
	Version int `json:"-" dynamo:"schema_version"`
	Username string `json:"username" dynamo:"username"`
	Email string `json:"email" dynamo:"email"`
	EmailVerified bool `json:"email_verified" dynamo:"email_verified"`
	PasswordHash []byte `json:"-" dynamo:"password_hash"`
    Created time.Time `json:"created" dynamo:"created_on"`
}

func NewUser(username string, email string, password string) *User {
    uid, _ := uuid.NewV4()
    passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)

    return &User {
        UID: uid,
        Version: UserVersion,
        Username: username,
        PasswordHash: passwordHash,
        Email: email,
        Created: time.Now(),
    }
}

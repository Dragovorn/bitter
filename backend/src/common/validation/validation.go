package validation

import (
    "github.com/gofrs/uuid"
    "github.com/guregu/dynamo"
    "main/src/common/aws/database"
    "math/rand"
    "time"
)

type Code struct {
    Code int `dynamo:"code"`
    UserId uuid.UUID `dynamo:"user_id"`
    Created time.Time `dynamo:"created_on"`
}

func NewCode(user uuid.UUID) *Code {
    rand.Seed(time.Now().UnixNano())

    return &Code{
        Code: rand.Int(), // TODO: Secure me!
        UserId: user,
        Created: time.Now(),
    }
}

func New(v interface {}) error {
    return Table().Put(database.Serialize(v)).Run()
}

func Table() dynamo.Table {
    return database.ValidationTable()
}

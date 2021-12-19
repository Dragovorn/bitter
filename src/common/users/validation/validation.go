package validation

import (
	"github.com/gofrs/uuid"
	"github.com/guregu/dynamo"
	"main/src/common"
	"main/src/common/aws/database"
	"math/rand"
	"time"
)

type Code struct {
	Code    string    `dynamo:"code"`
	UserId  uuid.UUID `dynamo:"user_id"`
	Created time.Time `dynamo:"created_on"`
}

func (c *Code) Formatted() string {
	return c.Code[0:5] + "-" + c.Code[5:10] // debt, codes could theoretically be longer, consider rewriting to be dynamic
}

func AsURL(code *Code) string { // TODO: Change the URL return to the baseURL validate
	return "" // TODO: Stub
}

func NewCode(user uuid.UUID) *Code {
	rand.Seed(time.Now().UnixNano())

	return &Code{
		Code:    common.RandNumber(10),
		UserId:  user,
		Created: time.Now(),
	}
}

func New(v interface{}) error {
	return Table().Put(database.Serialize(v)).Run()
}

func Table() dynamo.Table {
	return database.ValidationTable()
}

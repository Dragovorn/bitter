package common

type ConstantsProvider interface {
	UsersTable() string
    ValidationTable() string
	UsernameIndex() string
    UserIdIndex() string
	Email() string
	AWSRegion() string
}

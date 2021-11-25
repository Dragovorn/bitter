package common

type ConstantsProvider interface {
	UsersTable() string
	UsernameIndex() string
	Email() string
	AWSRegion() string
}

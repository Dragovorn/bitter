package common

var constants ConstantsProvider

func Constants() ConstantsProvider {
    return constants
}

func Init(provider ConstantsProvider) {
    constants = provider
}

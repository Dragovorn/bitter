package common

import (
	"math/rand"
	"strconv"
)

var constants ConstantsProvider

func RandNumber(length int) string {
	var resultStr string

	for x := 0; x < length; x++ {
		resultStr += strconv.Itoa(rand.Intn(9))
	}

	return resultStr
}

func Constants() ConstantsProvider {
	return constants
}

func Init(provider ConstantsProvider) {
	constants = provider
}

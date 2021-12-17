package common

import (
	"math/rand"
	"strconv"
)

var constants ConstantsProvider

func RandNumber(length int) int {
	var resultStr string

	for x := 0; x < length; x++ {
		resultStr += strconv.Itoa(rand.Intn(9))
	}

	result, _ := strconv.Atoi(resultStr)

	return result
}

func Constants() ConstantsProvider {
	return constants
}

func Init(provider ConstantsProvider) {
	constants = provider
}

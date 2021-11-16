package common

import "encoding/json"

func PrettyJsonString(i interface {}) string {
    result, _ := json.MarshalIndent(i, "", "    ")

    return string(result)
}
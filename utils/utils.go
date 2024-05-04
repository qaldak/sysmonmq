package utils

import (
	"encoding/json"

	logger "github.com/qaldak/SysMonMQ/logging"
)

func GenerateJson(info interface{}) (string, error) {
	logger.Debug("")

	// generate JSON
	jsonBytes, err := json.Marshal(info)
	if err != nil {
		logger.Error("Error in Marshaling: ", err)
		return "", err
	}

	logger.Debug(string(jsonBytes))
	return string(jsonBytes), nil
}

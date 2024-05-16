package utils

import (
	"encoding/json"
	"os"
	"strings"

	logger "github.com/qaldak/sysmonmq/internal/logging"
)

/*
Convert incoming data into JSON and return.
*/
func GenerateJson(data interface{}) (string, error) {
	logger.Debug("") // Todo: check

	// generate JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error("Error in Marshaling: ", err)
		return "", err
	}

	logger.Debug(string(jsonBytes))
	return string(jsonBytes), nil
}

/*
Retrieves the hostname of the current machine. The optional variadic argument CapType enum can have values Lower (convert to lowercase) or Upper (convert to uppercase).
*/
func GetHostname(c ...CapType) (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Error("Failed to determine hostname. Error: ", err)
	}

	if len(c) > 0 {
		switch c[0] {
		case Lower:
			hostname = strings.ToLower(hostname)
		case Upper:
			hostname = strings.ToUpper(hostname)
		default:
			// return hostname as is
		}
	}

	return hostname
}

// Capitalization type
type CapType int

const (
	AsIs CapType = iota
	Lower
	Upper
)

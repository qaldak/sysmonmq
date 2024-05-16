package utils

import (
	"testing"
)

func TestGenerateJson(t *testing.T) {
	// Call the InitLogger function
	GenerateJson("Foo")

	// Add your assertions here to verify that the logger was initialized correctly
	// For example, you can check if the logger variable is not nil

	// If the assertions fail, you can use the t.Errorf function to report the error
	// For example:
	// if logger == nil {
	//     t.Errorf("Expected logger to be initialized, but got nil")
	// }
}

// Add more test functions for the other functions in the logger.go file

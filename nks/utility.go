package nks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// PrettyPrint to break down objects
func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}

// GetEnvID grabs string from environment and converts to integer
func GetIDFromEnv(name string) (int, error) {
	var id int
	fmt.Sscanf(os.Getenv(name), "%d", &id)
	if id == 0 {
		return id, errors.New("Missing ID env in " + name)
	}
	return id, nil
}

// StringInSlice utlity function, like in_array, userful for validation of machine types
func StringInSlice(s string, list []string) bool {
	for _, item := range list {
		if s == item {
			return true
		}
	}
	return false
}

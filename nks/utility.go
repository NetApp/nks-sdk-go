package nks

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

//GetValueFromEnv grabs string from environment
func GetValueFromEnv(name string) (string, error) {
	content := os.Getenv(name)
	if len(content) == 0 {
		return "", errors.New("Empty content of env " + name)
	}
	return content, nil
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

//GetTicks gets the string representation of current ticks
func GetTicks() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func GetAbsPath(path string) (string, error) {
	usr, _ := user.Current()
	dir := usr.HomeDir

	if !filepath.IsAbs(path) {
		if path == "~" {
			// In case of "~", which won't be caught by the "else if"
			path = dir
		} else if strings.HasPrefix(path, "~/") {
			// Use strings.HasPrefix so we don't match paths like
			// "/something/~/something/"
			path = filepath.Join(dir, path[2:])
		}
	}

	return path, nil
}

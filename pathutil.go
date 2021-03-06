// Package pathutil contains utility functions for working paths
// It is also a wrapper for path & filepath which can be accessed directly
package pathutil

import (
	"os"
	pathlib "path"
)

// Check if a path exists regardless of whether it is a dir or file
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Expands '~/' prefix to the home directory
func Expand(path string) string {
	if len(path) < 2 {
		return path
	}
	if path[:2] == "~/" {
        home := os.Getenv("HOME")
		path = pathlib.Join(home, path[2:])
	}
	return path
}

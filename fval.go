// Package fval contains simple utility methods to validate file and directory
// existence
package fval

import (
	"fmt"
	"os"
)

// FileExists returns true if a file exists at path, false otherwise.
func FileExists(path string) bool {
	exists, f := exists(path)
	if exists && f != nil {
		if !f.IsDir() {
			return true
		}
	}
	return false
}

// DirExists returns true if a directory exists at path, false otherwise.
func DirExists(path string) bool {
	exists, d := exists(path)
	if exists && d != nil {
		if d.IsDir() {
			return true
		}
	}
	return false
}

// Exists returns true if either a file or directory exists at path,
// false otherwise.
func Exists(path string) bool {
	exists, _ := exists(path)
	return exists
}

// DirExistsOrCreate checks if a directory exists or path, if not it attempts
// to create it, and all parent directories, using os.MkdirAll.
// Returns ok if the file exists or if it could be created and error
// if something goes wrong when creating the directory.
func DirExistsOrCreate(path string, perm os.FileMode) (bool, error) {
	if DirExists(path) {
		return true, nil
	}

	err := os.MkdirAll(path, perm)
	if err != nil {
		return DirExists(path), err
	}

	return DirExists(path), nil
}

// DirPurgeAndCreate removes directory at path with all sub directories and files
// then creates a new empty directory at path.
// Returns ok if the directory is created. Returns error if no dir exists at path
// or if something goes wrong when deleting files or creating directory.
func DirPurgeAndCreate(path string, perm os.FileMode) (bool, error) {
	if !DirExists(path) {
		return false, fmt.Errorf("fval: trying to purge and create non existing dir")
	}

	err := os.RemoveAll(path)
	if err != nil {
		return false, err
	}

	return DirExistsOrCreate(path, perm)
}

// exists returns true if a file or directory exists at path, false otherwise.
func exists(path string) (bool, os.FileInfo) {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
	}
	return true, f
}

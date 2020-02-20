package fileio

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bugsnag/osext"
)

// GetContent returns the content of a given file. It first
// checks to see if the path is absolute. If not, then it
// checks for filepath in the current working directory.
// If still not found, it checks in the executable directory.
func GetContent(path string) (string, error) {
	var absPath string
	var exists bool

	// if absolute path, file is located
	if filepath.IsAbs(path) {
		absPath = path
		_, err := os.Stat(absPath)
		exists = err == nil
	}

	// look in current workding directory
	if !exists {
		var err error
		absPath, err = filepath.Abs(path) // creates current working directory
		if err != nil {
			_, err = os.Stat(absPath)
		}
		exists = err == nil
	}

	// it in executable directory
	if !exists {
		edir, ferr := osext.ExecutableFolder()
		if ferr == nil {
			absPath = filepath.Clean(edir + "/" + path)
			_, err := os.Stat(absPath)
			exists = err == nil
		}
	}

	if !exists {
		return "", errors.New("File not found in working or exec directory: " + path)
	}

	// read file and return content
	b, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ReadFile(adap, path string, param ...string) (string, error) {

	if adap == "disk" {
		return GetContent(path)
	}

	return "", errors.New("Unsupported adapter")
}

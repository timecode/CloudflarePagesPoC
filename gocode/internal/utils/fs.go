package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// MakeDirPath ...
func MakeDirPath(in string) (err error) {
	err = os.MkdirAll(filepath.Dir(in), 0770)
	return
}

// LoadFile ...
func LoadFile(in string) (out []byte, err error) {
	var file *os.File
	if file, err = os.Open(in); err != nil {
		return
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// SaveFile ...
func SaveFile(in string, b []byte) (err error) {
	if err = MakeDirPath(in); err != nil {
		return
	}
	err = ioutil.WriteFile(in, b, 0644)
	return
}

// FileExistsAtFilepath ...
func FileExistsAtFilepath(filepath string) (exists bool, err error) {
	_, err = os.Stat(filepath)
	if err == nil {
		exists = true
	} else if os.IsNotExist(err) {
	} else {
		// Schrödinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		err = fmt.Errorf("error unable to determine if file '%v' exists or not; please report to developer", filepath)
	}
	return
}

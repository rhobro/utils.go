package fileio

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var TempDir string

func initialize(dir, pattern string) {
	TempDir, _ = ioutil.TempDir(dir, pattern)
}

func TempPath(path string) (string, error) {
	// Make chain of paths
	pthSplit := strings.Split(path, string(filepath.Separator))
	totPath := make([]string, len(pthSplit)+1)
	totPath[0] = TempDir
	for i := 0; i < len(pthSplit); i++ {
		totPath[i+1] = pthSplit[i]
	}

	err := os.MkdirAll(filepath.Join(totPath[:len(totPath)-1]...), 0700)
	if err != nil {
		return "", fmt.Errorf("unable to create tmp path: %s", err)
	}

	return filepath.Join(totPath...), nil
}

package fileio

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var TmpDir string

func Init(inDir, name string) {
	var err error
	TmpDir, err = ioutil.TempDir(inDir, name)
	if err != nil {
		rd := bufio.NewScanner(os.Stdin)
		fmt.Println("Unable to create temp dir")
		fmt.Print("Enter new temp dir: ")
		rd.Scan()
		TmpDir = rd.Text()
	}
}

func TmpPath(path string) (string, error) {
	// Make chain of paths
	pthSplit := strings.Split(path, string(filepath.Separator))
	totPath := make([]string, len(pthSplit)+1)
	totPath[0] = TmpDir
	for i := 0; i < len(pthSplit); i++ {
		totPath[i+1] = pthSplit[i]
	}

	err := os.MkdirAll(filepath.Join(totPath[:len(totPath)-1]...), 0700)
	if err != nil {
		return "", fmt.Errorf("unable to create tmp path: %s", err)
	}

	return filepath.Join(totPath...), nil
}

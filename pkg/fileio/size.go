/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: size.go
 * Last Modified: 25/05/2021, 21:44
 */

package fileio

import (
	"os"
	"path/filepath"
)

func DirSize(path string) (size int64, err error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			// recurse
			sz, err := DirSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return 0, err
			}

			size += sz

		} else {
			// file
			info, err := entry.Info()
			if err != nil {
				return 0, err
			}

			size += info.Size()
		}
	}
	return
}

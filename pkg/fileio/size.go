/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: size.go
 * Last Modified: 25/05/2021, 21:16
 */

package fileio

import (
	"os"
)

func DirSize(path string) (size int64, err error) {
	entries, err := os.ReadDir(path)

	for _, entry := range entries {
		if entry.IsDir() {
			// recurse
			sz, err := DirSize(entry.Name())
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

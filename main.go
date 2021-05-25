/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: main.go
 * Last Modified: 25/05/2021, 21:12
 */

package main

import (
	"fmt"
	"github.com/rhobro/goutils/pkg/fileio"
)

func main() {
	fmt.Println(fileio.DirSize("/Users/robro/Desktop/media"))
}

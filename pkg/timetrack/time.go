/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: time.go
 * Last Modified: 15/01/2021, 08:30
 */

package timetrack

import (
	"fmt"
	"time"
)

var t = time.Now()
var totStart = time.Now()

func Init() {
	t = time.Now()
	totStart = time.Now()
}

func SinceLast(s string) {
	fmt.Printf("%20s : %v\n", s, time.Since(t))
	t = time.Now()
}

func Total(s string) {
	fmt.Printf("%20s : %v\n", s, time.Since(totStart))
}

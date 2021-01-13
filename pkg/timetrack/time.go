/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: time.go
 * Last Modified: 13/01/2021, 16:12
 */

package timetrack

import (
	"fmt"
	"sync"
	"time"
)

var t = time.Now()
var totStart = time.Now()
var initOnce sync.Once

func Init() {
	t = time.Now()
	totStart = time.Now()
}

func SinceLast(s string) {
	initOnce.Do(Init)
	fmt.Printf("%20s : %v\n", s, time.Since(t))
	t = time.Now()
}

func Total(s string) {
	fmt.Printf("%20s : %v\n", s, time.Since(totStart))
}

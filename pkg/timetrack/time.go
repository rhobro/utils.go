package timetrack

import (
	"fmt"
	"time"
)

var SinceLast = func(s string) {
	if !initd {
		Init()
		initd = true
	}
	fmt.Printf("%20s : %v\n", s, time.Since(t))
	t = time.Now()
}
var Total = func(s string) {
	fmt.Printf("%20s : %v\n", s, time.Since(totStart))
}
var initd bool
var t = time.Now()
var totStart = time.Now()

func Init() {
	t = time.Now()
	totStart = time.Now()
}

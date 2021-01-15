/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: func.go
 * Last Modified: 15/01/2021, 20:33
 */

package httputil

import (
	"fmt"
	"github.com/Bytesimal/goutils/pkg/util"
	"net/http"
	"strings"
	"time"
)

const defaultRetries = 10

func RQUntil(cli *http.Client, rq *http.Request) (*http.Response, error) {
	return RQUntilCustom(cli, rq, defaultRetries)
}

func RQUntilCustom(cli *http.Client, rq *http.Request, count int) (rsp *http.Response, err error) {
	err = fmt.Errorf("tmp")
	for i := 0; i < count; i++ {
		if err != nil {
			if i > 0 {
				// Random sleep in millisecond
				time.Sleep(time.Duration(util.Rand.Int63n(int64(time.Millisecond))))
			}
			rsp, err = cli.Do(rq)
			continue
		}
		break
	}
	return
}

const std = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.1 Safari/605.1.15"

const (
	macOS = iota
	windows
	iPod
	iPhone
	iPad
)
const nOS = 5

var iOSDevices = map[int]string{
	iPod:   "iPod",
	iPhone: "iPhone",
	iPad:   "iPad",
}

func RandUA() (ua string) {
	os := util.Rand.Intn(nOS)
	ua = "Mozilla/5.0 ("

	switch os {
	case macOS:
		ua += "Macintosh; Intel Mac OS X "
		// choose browser: 0 - safari, 1 - firefox, 2 - chrome
		br := util.Rand.Intn(3)
		if br == 1 {
			ua += macVs[util.Rand.Intn(len(macVs))] + "; rv:70.0) Gecko/20100101 Firefox/70.0"
			break
		}

		ua += strings.ReplaceAll(macVs[util.Rand.Intn(len(macVs))], ".", "_") + ") "
		ua += "AppleWebKit/" + webkitVs["mac"][util.Rand.Intn(len(webkitVs["mac"]))] + " (KHTML, like Gecko)"

	case windows:
		ua += "Windows NT 10.0; Win64; x64"
		// choose browser: firefox, safari, edge
		br := util.Rand.Intn(3)
		if br == 0 {
			ua += "; rv:70.0) Gecko/20100101 Firefox/70.0"
			break
		}

		safariV := webkitVs["mac"][util.Rand.Intn(len(webkitVs["mac"]))]
		ua += fmt.Sprintf(") AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s",
			safariV,
			chromeVs[util.Rand.Intn(len(chromeVs))],
			safariV)

		if br == 2 {
			ua += " Edge/" + edgeHTMLVs[util.Rand.Intn(len(edgeHTMLVs))]
		}

	default:
		// iOS devices
		ua += fmt.Sprintf("%s; CPU iPhone OS %s like Mac OS X) AppleWebKit/%s (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/%s",
			iOSDevices[os],
			iosVs[util.Rand.Intn(len(iosVs))],
			webkitVs["ios"][util.Rand.Intn(len(webkitVs["ios"]))],
			webkitVs["ios"][util.Rand.Intn(len(webkitVs["ios"]))])
	}

	return ua
}

/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: func.go
 * Last Modified: 28/01/2021, 19:45
 */

package httputil

import (
	"errors"
	"fmt"
	"github.com/Bytesimal/goutils/pkg/util"
	"math"
	"net/http"
	"strconv"
	"strings"
)

const defaultRetries = 10

func RQUntil(cli *http.Client, rq *http.Request) (*http.Response, error) {
	return RQUntilCustom(cli, rq, defaultRetries)
}

// Param count specifies the number of retries. If count == -1, then it will try infinitely
func RQUntilCustom(cli *http.Client, rq *http.Request, count int) (rsp *http.Response, err error) {
	var infinite bool
	if count == -1 {
		infinite = true
		count = 1
	}

	err = errors.New("")
	for i := 0; i < count; i++ {
		if err != nil {
			rsp, err = cli.Do(rq)

			// infinite
			if infinite {
				count += 1
			}

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

func IsValidIPv4(ip string) (valid bool) {
	valid = len(ip) < 16
	valid = strings.Count(ip, ".") == 3 && valid
	valid = !strings.ContainsAny(ip, "abcdefghijklmnopqrstuvwxyz") && valid

	for _, rawN := range strings.Split(ip, ".") {
		parsed, err := strconv.Atoi(rawN)
		valid = parsed <= math.MaxUint8 && valid
		valid = err == nil && valid
	}
	return
}

/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: func.go
 * Last Modified: 15/02/2021, 10:38
 */

package httputil

import (
	"errors"
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

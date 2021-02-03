/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: sentry.go
 * Last Modified: 03/02/2021, 20:43
 */

package sentree

import (
	"github.com/getsentry/sentry-go"
	"log"
)

var C *sentry.Client

func Init(opt sentry.ClientOptions, verbose bool) {
	var err error
	C, err = sentry.NewClient(opt)
	if err != nil {
		log.Fatalf("can't connect to sentree: %s", err)
	}

	if verbose {
		log.Print("{sentree} connected")
	}
}

/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: sentry.go
 * Last Modified: 24/05/2021, 11:32
 */

package sentree

import (
	"github.com/getsentry/sentry-go"
	"log"
)

var C *sentry.Client

func Init(opt sentry.ClientOptions, verbose bool) error {
	var err error
	C, err = sentry.NewClient(opt)
	if err != nil {
		return err
	}

	if verbose {
		log.Print("{sentry} connected")
	}
	return nil
}

func LogCaptureErr(err error) {
	if C == nil {
		panic("Sentry client is nil")
	}

	C.CaptureException(err, nil, nil)
	log.Print(err)
}

func FatalCaptureErr(err error) {
	if C == nil {
		panic("Sentry client is nil")
	}

	C.CaptureException(err, nil, nil)
	log.Fatal(err)
}

/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: utils.go
 * File Name: sentry.go
 * Last Modified: 28/08/2021, 13:14
 */

package sentree

import (
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

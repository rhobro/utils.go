/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: configcat.go
 * Last Modified: 02/02/2021, 15:13
 */

package cfgcat

import (
	"github.com/configcat/go-sdk/v7"
	"log"
)

var C *configcat.Client

func Init(key string, verbose bool) {
	C = configcat.NewClient(key)
	if verbose {
		log.Print("{configcat} connected")
	}
}

func InitCustom(cfg configcat.Config, verbose bool) {
	C = configcat.NewCustomClient(cfg)
	if verbose {
		log.Print("{configcat} connected")
	}
}

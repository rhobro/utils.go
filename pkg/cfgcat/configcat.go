/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: configcat.go
 * Last Modified: 02/02/2021, 09:57
 */

package cfgcat

import (
	"github.com/configcat/go-sdk/v7"
)

var C *configcat.Client

func Init(key string) {
	C = configcat.NewClient(key)
}

func InitCustom(cfg configcat.Config) {
	C = configcat.NewCustomClient(cfg)
}

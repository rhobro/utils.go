/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: util.go
 * Last Modified: 27/03/2021, 21:16
 */

package util

import (
	"errors"
	"log"

	"github.com/pariz/gountries"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var countries = gountries.New()

func FindCountryByISO(iso string) (gountries.Country, error) {
	for k, v := range countries.FindAllCountries() {
		if iso == k {
			return v, nil
		}
	}
	return gountries.Country{}, errors.New("invalid country ISO")
}

func StrIndex(t string, coll []string) int {
	for i, item := range coll {
		if item == t {
			return i
		}
	}
	return -1
}

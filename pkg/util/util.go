/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: goutils
 * File Name: util.go
 * Last Modified: 27/03/2021, 21:06
 */

package util

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/pariz/gountries"
)

func init() {
	now := time.Now().UnixNano()
	r = rand.New(rand.NewSource(now))
}

var mtx sync.Mutex
var r *rand.Rand

func Rand() *rand.Rand {
	mtx.Lock()
	defer mtx.Unlock()
	return r
}

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

package util

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/pariz/gountries"
)

func init() {
	now := time.Now().UnixNano()
	Rand = rand.New(rand.NewSource(now))
}

var Rand *rand.Rand

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

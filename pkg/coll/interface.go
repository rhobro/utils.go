/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: utils.go
 * File Name: interface.go
 * Last Modified: 28/08/2021, 13:14
 */

package coll

func Index(s []interface{}, x interface{}) int {
	for i, e := range s {
		if x == e {
			return i
		}
	}
	return -1
}

func Contains(s []interface{}, x interface{}) bool {
	return Index(s, x) > -1
}

func Remove(s []interface{}, i int) []interface{} {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

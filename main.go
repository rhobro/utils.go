/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: utils.go
 * File Name: main.go
 * Last Modified: 07/08/2021, 19:45
 */

package main

import "github.com/rhobro/utils.go/pkg/services/mongrel"

func main() {
	mongrel.Init("mongodb+srv://<username>:<password>@mongo.g15ne.mongodb.net/?retryWrites=true&w=majority")
}

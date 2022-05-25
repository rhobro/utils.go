/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: utils.go
 * File Name: init.go
 * Last Modified: 07/08/2021, 19:47
 */

package mongrel

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var C *mongo.Client

var DB *mongo.Database

func Init(url, db string) error {
	// connect
	var err error
	C, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	DB = C.Database(db, options.Database())

	return nil
}

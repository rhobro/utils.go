package mongrel

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var C *mongo.Client

func Init(url string) error {
	var err error
	C, err = mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	return nil
}

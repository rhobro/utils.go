package mongrel

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var C *mongo.Client

func Init(url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	var err error
	C, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	return nil
}

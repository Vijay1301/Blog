package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateNewConnection(mongoConfig *Mongo) (*mongo.Client, error) {
	opts := options.Client()
	opts.ApplyURI(mongoConfig.Url)
	if !mongoConfig.SkipAuth {
		creds := options.Credential{
			Username: mongoConfig.User,
			Password: mongoConfig.Password,
		}
		opts.SetAuth(creds)
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.PrimaryPreferred())
	if err != nil {
		return nil, err
	}

	return client, nil
}

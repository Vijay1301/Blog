package account

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DAO struct {
	db *mongo.Database
}

func NewDAO(db *mongo.Database) *DAO {
	return &DAO{
		db: db,
	}
}

func (d *DAO) CreateAccount(ctx context.Context, user AccountDao) error {
	usersColl := d.db.Collection("accounts")
	query := bson.M{
		"email":     user.Email,
		"accountId": user.AccountId,
		"userId":    user.UserId,
	}

	opts := options.Update().SetUpsert(true)
	_, err := usersColl.UpdateOne(ctx, query, bson.M{"$set": user}, opts)

	if err != nil {
		return err
	}
	return nil
}

func (d *DAO) FindAccount(ctx context.Context, email string) (*AccountDao, error) {
	usersColl := d.db.Collection("accounts")
	query := bson.M{
		"email": email,
	}
	var user AccountDao
	err := usersColl.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &user, err
}

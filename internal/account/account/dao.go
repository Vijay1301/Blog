package account

import (
	"context"
	"errors"

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

	accountColl := d.db.Collection("accounts")
	query := bson.M{
		"email": user.Email,
		"id":    user.Id,
	}

	opts := options.Update().SetUpsert(true)
	_, err := accountColl.UpdateOne(ctx, query, bson.M{"$set": user}, opts)

	if err != nil {
		return err
	}
	return nil
}

func (d *DAO) FindAccount(ctx context.Context, email string) (*AccountDao, error) {

	accountColl := d.db.Collection("accounts")

	query := bson.M{
		"email": email,
	}
	var user AccountDao
	err := accountColl.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, err
	}
	return &user, err
}

func (d *DAO) GetAccountById(ctx context.Context, Id string) (*Account, error) {

	accountColl := d.db.Collection("accounts")

	quey := bson.M{
		"id": Id,
	}

	var account Account

	err := accountColl.FindOne(ctx, quey).Decode(&account)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	return &account, nil
}

func (d *DAO) GetAccountForUpdate(ctx context.Context, Id string) (*AccountDao, error) {

	accountColl := d.db.Collection("accounts")

	quey := bson.M{
		"id": Id,
	}

	var account AccountDao

	err := accountColl.FindOne(ctx, quey).Decode(&account)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	return &account, nil
}


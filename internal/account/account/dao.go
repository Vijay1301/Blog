package account

import "go.mongodb.org/mongo-driver/mongo"

type DAO struct {
	db *mongo.Database
}

func NewDAO(db *mongo.Database) *DAO {
	return &DAO{
		db: db,
	}
}

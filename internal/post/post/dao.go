package post

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

func (d *DAO) CreatePost(ctx context.Context, req BlogPostDao) error {

	postColl := d.db.Collection("posts")

	query := bson.M{
		"account_id": req.AccountId,
		"id":         req.ID,
	}

	opts := options.Update().SetUpsert(true)
	_, err := postColl.UpdateOne(ctx, query, bson.M{"$set": req}, opts)

	if err != nil {
		return err
	}
	return nil
}

func (d *DAO) GetPostById(ctx context.Context, Id string, postId string) (*GetPost, error) {

	postColl := d.db.Collection("posts")

	query := bson.M{
		"account_id": Id,
		"id":         postId,
	}

	var post GetPost

	err := postColl.FindOne(ctx, query).Decode(&post)

	if err != nil {
		return nil, err
	}

	return &post, nil

}

func (d *DAO) GetAllPost(ctx context.Context, Id string) ([]GetPost, int64, error) {

	postColl := d.db.Collection("posts")

	query := bson.M{
		"account_id": Id,
	}

	var post []GetPost

	cursor, err := postColl.Find(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Decode all posts
	if err := cursor.All(ctx, &post); err != nil {
		return nil, 0, err
	}

	count, err := postColl.CountDocuments(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	return post, count, nil

}

func (d *DAO) DeletePost(ctx context.Context, Id string, postId string) error {

	postColl := d.db.Collection("posts")

	query := bson.M{
		"account_id": Id,
		"id":         postId,
	}

	_, err := postColl.DeleteOne(ctx, query)

	if err != nil {
		return err
	}

	return nil

}

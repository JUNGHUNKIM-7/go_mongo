package repository

import (
	"context"
	"errors"
	"os"

	"example.com/main/src/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connnect() error {
	c := &types.Config{
		Uri:      os.Getenv("URI"),
		DbName:   os.Getenv("DB"),
		CollName: os.Getenv("COLL"),
	}
	if c.CollName == "" || c.DbName == "" || c.Uri == "" {
		panic(".env vars are nil")
	}

	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.Uri))
	if err != nil {
		return errors.New("client err")
	}

	types.MongoClient = &types.Client{C: cli}

	db := cli.Database(c.DbName)
	coll := db.Collection(c.CollName)
	types.MongoConfigContainer = &types.MongoConfig{
		Db:   db,
		Coll: coll,
	}
	if types.MongoConfigContainer.Db == nil || types.MongoConfigContainer.Coll == nil {
		panic("Db or Coll is nil")
	}

	return nil
}

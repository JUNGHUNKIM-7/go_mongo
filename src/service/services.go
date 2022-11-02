package service

import (
	"context"
	"log"

	"example.com/main/src/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct{}

func (s Service) Coll() *mongo.Collection {
	container := types.MongoConfigContainer
	coll := container.Coll
	return coll
}

func (s Service) Find() ([]types.Pokemon, error) {
	coll := s.Coll()
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var result types.Pokemon
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		} else {
			types.ModelSli = append(types.ModelSli, result)
		}
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}

	return types.ModelSli, nil
}

func (s Service) Post() {

}

func (s Service) FindOne(id string) {
	// coll := s.getColl()
}

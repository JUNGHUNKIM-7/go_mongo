package types

import "go.mongodb.org/mongo-driver/mongo"

type Config struct {
	Uri, DbName, CollName string
}

type Client struct {
	C *mongo.Client
}

type MongoConfig struct {
	Db   *mongo.Database
	Coll *mongo.Collection
}

type Pokemon struct {
	Id              int      `bson:"id,omitempty"`
	Name            string   `bson:"name,omitempty"`
	Type            []string `bson:"type,omitempty"`
	Hp              int      `bson:"hp,omitempty"`
	Attack          int      `bson:"attack,omitempty"`
	Defense         int      `bson:"defense,omitempty"`
	Special_attack  int      `bson:"special_attack,omitempty"`
	Special_defense int      `bson:"special_defense,omitempty"`
	Speed           int      `bson:"speed,omitempte"`
	Upsert          bool     `bson:"upsert,omitempty"`
}

var ModelSli []Pokemon
var MongoClient *Client
var MongoConfigContainer *MongoConfig

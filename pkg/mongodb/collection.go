package mongodb

import "go.mongodb.org/mongo-driver/mongo"

func User() *mongo.Collection {
	return Db.Collection("user")
}

func Poll() *mongo.Collection {
	return Db.Collection("polls")
}

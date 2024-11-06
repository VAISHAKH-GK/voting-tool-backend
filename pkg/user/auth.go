package user

import (
	"context"

	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func (user *User) InsertUser() error {
	var _, err = mongodb.User().InsertOne(context.Background(), user)
	return err
}

func GetUser(username string) User {
	var user User
	var result = mongodb.User().FindOne(context.Background(), bson.M{"username": username})
	result.Decode(&user)
	return user
}

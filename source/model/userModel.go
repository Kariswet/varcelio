package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserModel struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
	Status   bool   `bson:"status"`
	Region   string `bson:"region"`
}

type User_Search struct {
	//? Regex
	Search string `json:"search"`

	Request
}

func (o *User_Search) HandleFilter(listFilterAnd *[]bson.M) {
	if search := o.Search; search != "" {
		*listFilterAnd = append(*listFilterAnd, bson.M{"username": primitive.Regex{Pattern: search, Options: "i"}})
	}
}

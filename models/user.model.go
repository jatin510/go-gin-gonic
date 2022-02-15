package models

import (
	"fmt"
	"gin-gonic-api/db"
	"gin-gonic-api/forms"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

func (h User) GetByID(id string) (*User, error) {
	db := db.GetDB()
	fmt.Print("db", db)

	// params := &dynamodb.GetItemInput{
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"user_id": {
	// 			S: aws.String(id),
	// 		},
	// 	},
	// 	TableName:      aws.String("TableUsers"),
	// 	ConsistentRead: aws.Bool(true),
	// }
	// resp, err := db.GetItem(params)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	var user *User
	// if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return user, nil
}

func (h User) Create(userSignUpForm forms.UserSignUp) (*User, error) {
	db := db.GetDB()
	fmt.Print("db", db)

	var user *User

	return user, nil
}

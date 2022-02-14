package repository

import (
	"gin-gonic-api/db"
	"gin-gonic-api/forms"
	"gin-gonic-api/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserEntity IUser

type userEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

type IUser interface {
	getAll() ([]models.User, int, error)
	getById(id string) (*models.User, int, error)
	CreateOne(userForm forms.UserSignUp) (*models.User, int, error)
}

// func (u User) GetByID(id string) (*User, error) {
// 	var user *User
// 	var err error

// 	return user, err
// }

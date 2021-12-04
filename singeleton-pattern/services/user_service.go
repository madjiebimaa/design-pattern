package services

import (
	"go-design-pattern/singeleton-pattern/apps"
	"go-design-pattern/singeleton-pattern/models"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Create(user models.User) {
	// create user with one connection
	db := apps.NewDatabaseHelper().GetConnection()
	db.Exec("INSERT INTO users ...")
}

func (s *UserService) Update(user models.User) {
	// update user with one connection
	db := apps.NewDatabaseHelper().GetConnection()
	db.Exec("SET FROM users ...")
}

func (s *UserService) FindAll() []models.User {
	// update user wit hone connection
	db := apps.NewDatabaseHelper().GetConnection()
	db.Query("SELECT * FROM users")

	return []models.User{}
}

package user

import "demo_project/database"

type UserRepositories struct{}

var db = database.GetInstance().GetDB()

func (ur *UserRepositories) Create(data interface{}) error {
	err := db.Create(data).Error
	return err
}

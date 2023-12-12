package managerCV

import "demo_project/database"

type CVRepositories struct{}

var db = database.GetInstance().GetDB()

func (cv *CVRepositories) Create(data interface{}) error {
	err := db.Create(data).Error
	return err
}

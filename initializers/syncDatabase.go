package initializers

import (
	"demo_project/database"
	"demo_project/model"
)

func SyncDatabase() {
	database.GetInstance().GetDB().AutoMigrate(&model.User{})
}

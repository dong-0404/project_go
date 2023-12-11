package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type Database struct {
	db *gorm.DB
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
		instance = &Database{}
		instance.connectDB()
	})
	return instance
}
func (d *Database) connectDB() {
	dsn := "root:shin2004@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	d.db = db
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

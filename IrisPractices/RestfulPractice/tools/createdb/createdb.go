package main

import (
	"github.com/jinzhu/gorm"
	"github.com/Humbertzhang/goweb_practice/IrisPractices/RestfulPractice/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	Db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Post{})
}

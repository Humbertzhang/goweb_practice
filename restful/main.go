package main

import (
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Db *gorm.DB

type User struct {
	ID               int
	Username         string `gorm:"unique"`
	Password_hash    string
}

func init() {
	var err error
	Db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&User{})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signin", signin)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
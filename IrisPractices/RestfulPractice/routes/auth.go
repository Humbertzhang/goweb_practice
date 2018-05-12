package routes

import (
	"github.com/kataras/iris"
	"github.com/Humbertzhang/IrisPractices/Practice1/models"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Post{})
}


func Register(ctx iris.Context) {
	var u models.User
	err := ctx.ReadJSON(&u)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	Db.Create(&u)
	ctx.JSON(map[string]int {
		"uid": u.ID,
	})
}

func Login(ctx iris.Context) {
	var u models.User
	var udb models.User

	err := ctx.ReadJSON(&u)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	if Db.Where("username = $1", u.Username).First(&udb).RecordNotFound() {
		ctx.StatusCode(404)
		return
	}
	if u.PassWord != udb.PassWord {
		ctx.StatusCode(401)
		return
	}

	ctx.JSON(map[string]string {
		"token": strconv.Itoa(udb.ID) + udb.Username,
	})
}
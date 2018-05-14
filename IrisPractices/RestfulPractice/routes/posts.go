package routes

import (
	"github.com/kataras/iris"
	"github.com/Humbertzhang/goweb_practice/IrisPractices/RestfulPractice/models"
	"strconv"
)

func PostArticle(ctx iris.Context) {
	token := ctx.GetHeader("token")

	var p models.Post
	err := ctx.ReadJSON(&p)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	var u models.User
	if Db.Where("id = $1", p.UserID).First(&u).RecordNotFound() {
		ctx.JSON(map[string]string {
			"msg": "user not found",
		})
		return
	}

	check := strconv.Itoa(u.ID) + u.Username
	if check != token {
		ctx.JSON(map[string]string {
			"msg": "token invalid",
		})
		return
	}

	Db.Create(&p)
	ctx.JSON(map[string]int {
		"postid": p.ID,
	})
}

func GetOneArticle(ctx iris.Context) {
	var p models.Post

	pid, _ := ctx.Params().GetInt("postid")

	if Db.Where("id = $1", pid).First(&p).RecordNotFound() {
		ctx.JSON(map[string]string {
			"msg": "post not found",
		})
	}
	ctx.JSON(p)
}

func GetMyArticles(ctx iris.Context) {
	uid, _ := ctx.Params().GetInt64("uid")
	var u models.User
	if Db.Where("id = $1", uid).First(&u).RecordNotFound() {
		ctx.JSON(map[string]string {
			"msg": "user not found",
		})
	}
	var posts[] models.Post
	Db.Model(&u).Related(&posts)
	ctx.JSON(posts)
}

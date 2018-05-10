package routes

import (
	"github.com/kataras/iris"
	"github.com/Humbertzhang/IrisPractices/Practice1/models"
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
	if string(u.ID) + u.Username != token {
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

	pid := ctx.Params().Get("postid")

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
	ctx.JSON(u.Posts)
}
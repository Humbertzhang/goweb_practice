package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Humbertzhang/goweb_practice/Configtest/router/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Secure)
	g.Use(middleware.Options)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "No Such API")
	})
	return g
}
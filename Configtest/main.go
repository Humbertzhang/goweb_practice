package main

import (
	"github.com/Humbertzhang/goweb_practice/Configtest/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/Humbertzhang/goweb_practice/Configtest/router"
	"log"
	"net/http"
)

func main(){
	if err := config.Init(); err != nil {
		panic(err)
	}
	// viper 从config file 中读取
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	log.Printf("Start to listening the incoming requests on host:", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}
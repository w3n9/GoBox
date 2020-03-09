package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var r *gin.Engine

func init(){
	r=gin.Default()
	r.GET("/ping",func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,"pong")
	})
}

func Get() *gin.Engine{
	return r
}
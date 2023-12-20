package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joaquinxtomas/gin-gonic/controller"
	"github.com/joaquinxtomas/gin-gonic/middlewares"
	"github.com/joaquinxtomas/gin-gonic/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	router := gin.New()

	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	router.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	router.Run()
}

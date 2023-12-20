package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joaquinxtomas/gin-gonic/entity"
	"github.com/joaquinxtomas/gin-gonic/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(videoService service.VideoService) VideoController {
	return &controller{
		service: videoService,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

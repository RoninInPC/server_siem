package api

import (
	"github.com/gin-gonic/gin"
	"server_siem/command"
)

type Api struct {
	engine *gin.Engine
}

func InitApi() Api {
	return Api{gin.Default()}
}

func (api Api) Get(path string, action command.Action) {
	api.engine.GET(path, func(c *gin.Context) {
		action.Action(c)
	})
}

func (api Api) Post(path string, action command.Action) {
	api.engine.POST(path, func(c *gin.Context) {
		action.Action(c)
	})
}

func (api Api) Put(path string, action command.Action) {
	api.engine.PUT(path, func(c *gin.Context) {
		action.Action(c)
	})
}

func (api Api) Delete(path string, action command.Action) {
	api.engine.DELETE(path, func(c *gin.Context) {
		action.Action(c)
	})
}

func (api Api) Patch(path string, action command.Action) {
	api.engine.PATCH(path, func(c *gin.Context) {
		action.Action(c)
	})
}

func (api Api) Head(path string, action command.Action) {
	api.engine.HEAD(path, func(c *gin.Context) {
		action.Action(c)
	})
}

func (api Api) Options(path string, action command.Action) {
	api.engine.OPTIONS(path, func(c *gin.Context) {
		action.Action(c)
	})
}

func (api Api) Run(address string) {
	api.engine.Run(address)
}

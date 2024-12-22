package command

import "github.com/gin-gonic/gin"

type Action interface {
	Action(*gin.Context)
}

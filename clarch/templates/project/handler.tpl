package project

import (
	"strconv"

	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project/logger"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerProject struct{}

func (this *HandlerProject) ParseID(c *gin.Context) int {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		return id
	}

	return -1
}

func (this *HandlerProject) ParseInt(c *gin.Context, key string) int {
	if id, err := strconv.Atoi(c.Param(key)); err == nil {
		return id
	}

	return -1
}

func (this *HandlerProject) ValidateBindJSON(c *gin.Context, o interface{}) error {
	if err := c.BindJSON(&o); err != nil {
		logger.Logger.Error(err.Error())
		return err
	}

	if _, err := govalidator.ValidateStruct(o); err != nil {
		logger.Logger.Error(err.Error())
		return err
	}

	return nil
}

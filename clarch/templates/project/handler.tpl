package project

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerProject struct{}

func (this *HandlerProject) ValidateBindJSON(c *gin.Context, o interface{}) error {
	if err := c.BindJSON(&o); err != nil {
		return err
	}

	if _, err := govalidator.ValidateStruct(o); err != nil {
		return err
	}

	return nil
}

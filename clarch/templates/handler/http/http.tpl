package http

import (
	"net/http"

	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}/usecase"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project"
	"github.com/gin-gonic/gin"
)


type Http{{.CamelPkg}}Handler struct {
	project.HandlerProject
	{{.Pkg}}UC usecase.{{.CamelPkg}}Usecase
}

func (this *Http{{.CamelPkg}}Handler) List(c *gin.Context) {
	where := &model.{{.CamelPkg}}{}
	list, err := this.{{.Pkg}}UC.List(where)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (this *Http{{.CamelPkg}}Handler) Show(c *gin.Context) {
	id := this.ParseId(c)
	where := &model.{{.CamelPkg}}{Id: id}
	item, err := this.{{.Pkg}}UC.Show(where)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (this *Http{{.CamelPkg}}Handler) Store(c *gin.Context) {
	params := &StoreParams{}
	if err := this.ValidateBindJSON(c, params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	v := &model.{{.CamelPkg}}{}
	res, err := this.{{.Pkg}}UC.Store(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (this *Http{{.CamelPkg}}Handler) Update(c *gin.Context) {
	params := &UpdateParams{}
	if err := this.ValidateBindJSON(c, params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id := this.ParseID(c)
	v := &model.{{.CamelPkg}}{
		Id: id,
	}

	res, err := this.{{.Pkg}}UC.Update(v)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (this *Http{{.CamelPkg}}Handler) Delete(c *gin.Context) {
	id := this.ParseID(c)

	where := &model.{{.CamelPkg}}{
		Id:     id,
	}
	if err := this.{{.Pkg}}UC.Delete(where); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func New{{.CamelPkg}}HttpHandler(g *gin.Engine, uc usecase.{{.CamelPkg}}Usecase) {
	h := &Http{{.CamelPkg}}Handler{
		{{.Pkg}}UC: uc,
	}
	v1 := g.Group("/api/v1")
	v1.GET("/{{.Pkg}}s", h.List)
	v1.GET("/{{.Pkg}}s/:id", h.Show)
	v1.POST("/{{.Pkg}}s", h.Store)
	v1.PUT("/{{.Pkg}}s/:id", h.Update)
	v1.DELETE("/{{.Pkg}}s/:id", h.Delete)
}

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

func (this *Http{{.CamelPkg}}Handler) FindAll(c *gin.Context) {
	list, err := this.{{.Pkg}}UC.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (this *Http{{.CamelPkg}}Handler) FindById(c *gin.Context) {
	id := c.GetInt("id")
	item, err := this.{{.Pkg}}UC.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (this *Http{{.CamelPkg}}Handler) Store(c *gin.Context) {
	params := &{{.Pkg}}.{{.CamelPkg}}{}
	if err := this.ValidateBindJSON(c, &params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := this.{{.Pkg}}UC.Store(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (this *Http{{.CamelPkg}}Handler) Update(c *gin.Context) {
	id := this.ParseID(c)
	item, err := this.{{.Pkg}}UC.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := this.ValidateBindJSON(c, &item); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := this.{{.Pkg}}UC.Update(id, item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (this *Http{{.CamelPkg}}Handler) Delete(c *gin.Context) {
	id := this.ParseID(c)
	if err := this.{{.Pkg}}UC.Delete(id); err != nil {
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
	v1.GET("/{{.Pkg}}s", h.FindAll)
	v1.GET("/{{.Pkg}}s/:id", h.FindById)
	v1.POST("/{{.Pkg}}s", h.Store)
	v1.PUT("/{{.Pkg}}s/:id", h.Update)
	v1.DELETE("/{{.Pkg}}s/:id", h.Delete)
}

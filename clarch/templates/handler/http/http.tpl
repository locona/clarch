package http

import (
	"net/http"

	"github.com/{{.GithubUser}}/{{.GithubRepository}}/{{.Pkg}}"
	"github.com/{{.GithubUser}}/{{.GithubRepository}}/{{.Pkg}}/usecase"
	"github.com/{{.GithubUser}}/{{.GithubRepository}}/project"
	"github.com/gin-gonic/gin"
)


type Http{{.CamelPkg}}Handler struct {
	project.HandlerProject
	{{.Pkg}}UC usecase.{{.CamelPkg}}UseCase
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
	var o {{.Pkg}}.{{.CamelPkg}}
	if err := c.BindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := this.{{.Pkg}}UC.Store(&o)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (this *Http{{.CamelPkg}}Handler) Update(c *gin.Context) {
	var o {{.Pkg}}.{{.CamelPkg}}
	if err := c.BindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := this.{{.Pkg}}UC.Update(&o)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (this *Http{{.CamelPkg}}Handler) Delete(c *gin.Context) {
	id := c.GetInt("id")
	if err := this.{{.Pkg}}UC.Delete(id);err!=nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
}

func New{{.CamelPkg}}HttpHandler(g *gin.Engine, uc usecase.{{.CamelPkg}}Usecase) {
	h := &Http{{.CamelPkg}}Handler{
		{{.Pkg}}UC: uc,
	}
	g.GET("/{{.Pkg}}s", h.FindAll)
	g.GET("/{{.Pkg}}s/:id", h.FindById)
	g.POST("/{{.Pkg}}s", h.Store)
	g.PUT("/{{.Pkg}}s/:id", h.Update)
	g.DELETE("/{{.Pkg}}s/:id", h.Delete)
}

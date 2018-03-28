package repository

import "github.com/{{.GithubUser}}/{{.GithubRepository}}/{{.Pkg}}"

type {{.CamelPkg}}Repository interface {
	FindAll() ([]*{{.Pkg}}.{{.CamelPkg}}, error)
	FindById(id int) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Store(*{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Update(*{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Delete(id int) error
}

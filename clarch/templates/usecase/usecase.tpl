package usecase

import (
	"errors"
	"github.com/{{.GithubUser}}/{{.GithubRepository}}/{{.Pkg}}"
	"github.com/{{.GithubUser}}/{{.GithubRepository}}/{{.Pkg}}/repository"
)

type {{.CamelPkg}}Usecase interface {
	FindAll() ([]*{{.Pkg}}.{{.CamelPkg}}, error)
	FindById(id int) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Store(*{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Update(*{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Delete(id int) error
}

type {{.Pkg}}Usecase struct {
	{{.Pkg}}Repo repository.{{.CamelPkg}}Repository
}

func New{{.CamelPkg}}Usecase(r repository.{{.CamelPkg}}Repository) {{.CamelPkg}}Usecase {
	return &{{.Pkg}}Usecase{
		{{.Pkg}}Repo: r,
	}
}

func (this *{{.Pkg}}Usecase) FindAll() (*{{.Pkg}}.{{.CamelPkg}}, error) {
	res, err := this.{{.Pkg}}Repo.FindAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (this *{{.Pkg}}Usecase) FindById(id int) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	res, err := this.{{.Pkg}}Repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (this *{{.Pkg}}Usecase) Store(value *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	return this.{{.Pkg}}Repo.Store(value)
}

func (this *{{.Pkg}}Usecase) Update(id int, value *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
        existed, err := this.{{.Pkg}}Repo.FindById(id)
	if existed == nil {
		return nil, errors.New("Not found")
	}
	if err != nil {
		return nil, err
	}
	return this.{{.Pkg}}Repo.Update(id, value)
}

func (this *{{.Pkg}}Usecase) Delete(id int) error {
	existed, _ := this.{{.Pkg}}Repo.FindById(id)
	if existed == nil {
		return errors.New("Not found")
	}

	err := this.{{.Pkg}}Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

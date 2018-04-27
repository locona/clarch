package usecase

import (
	"errors"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/model"
	{{.Pkg}}Repo "github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}/repository"
)

type {{.CamelPkg}}Usecase interface {
	List(*model.{{.CamelPkg}}) ([]*model.{{.CamelPkg}}, error)
	Show(*model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error)
	Store(*model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error)
	Update(*model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error)
	Delete(*model.{{.CamelPkg}}) error
}

type {{.Pkg}}Usecase struct {
	{{.Pkg}}Repo {{.Pkg}}Repo.{{.CamelPkg}}Repository
}

func New{{.CamelPkg}}Usecase(r {{.Pkg}}Repo.{{.CamelPkg}}Repository) {{.CamelPkg}}Usecase {
	return &{{.Pkg}}Usecase{
		{{.Pkg}}Repo: r,
	}
}

func (this *{{.Pkg}}Usecase) List(where *model.{{.CamelPkg}}) ([]*model.{{.CamelPkg}}, error) {
	res, err := this.{{.Pkg}}Repo.FindBy(where)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (this *{{.Pkg}}Usecase) Show(where *model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error) {
	res, err := this.{{.Pkg}}Repo.FirstBy(where)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (this *{{.Pkg}}Usecase) Store(value *model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error) {
	return this.{{.Pkg}}Repo.Store(value)
}

func (this *{{.Pkg}}Usecase) Update(value *model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error) {
	where := &model.{{.CamelPkg}} {
		ID: value.ID,
	}
        if _, err := this.{{.Pkg}}Repo.FirstBy(where);err == nil {
		err = errors.New("Not found")
		return nil, err
	}

	res, err := this.{{.Pkg}}Repo.Update(value)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (this *{{.Pkg}}Usecase) Delete(where *model.{{.CamelPkg}}) error {
	existed, err := this.{{.Pkg}}Repo.FirstBy(where)
	if err != nil {
		return errors.New("Not found")
	}

	if err := this.{{.Pkg}}Repo.Delete(existed);err != nil{
		return err
	}

	return nil
}

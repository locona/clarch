package repository

import (
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project"
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/model"
	"github.com/jinzhu/gorm"
)

type {{.CamelPkg}}Repository interface {
	FindBy(*model.{{.CamelPkg}}) ([]*model.{{.CamelPkg}}, error)
	FirstBy(*model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error)
	Store(*model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error)
	Update(*model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error)
	Delete(*model.{{.CamelPkg}}) error
}

type {{.Pkg}}Repository struct {
	project.RepositoryProject
	DB *gorm.DB
}

func New{{.CamelPkg}}Repository(DB *gorm.DB) {{.CamelPkg}}Repository {
	return &{{.Pkg}}Repository{DB: DB}
}

func (this *{{.Pkg}}Repository) FindBy(where *model.{{.CamelPkg}}) ([]*model.{{.CamelPkg}}, error) {
	item := make([]*model.{{.CamelPkg}}, 0)
	if err := this.DB.Where(where).Find(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (this *{{.Pkg}}Repository) FirstBy(where *model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error) {
	item := &model.{{.CamelPkg}}{}
	if err := this.DB.Where(where).First(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (this *{{.Pkg}}Repository) Store(value *model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error) {
	if err := this.Transaction(this.DB, func(tx *gorm.DB) error {
		return tx.Create(value).Error
	}); err != nil {
		return nil, err
	}

	return value, nil
}

func (this *{{.Pkg}}Repository) Update(value *model.{{.CamelPkg}}) (*model.{{.CamelPkg}}, error) {
	if err := this.Transaction(this.DB, func(tx *gorm.DB) error {
		return tx.Save(value).Error
	}); err != nil {
		return nil, err
	}

	return value, nil
}

func (this *{{.Pkg}}Repository) Delete(where *model.{{.CamelPkg}}) error {
	item := &model.{{.CamelPkg}}{}
	if err := this.DB.Where(where).Delete(item).Error; err != nil {
		return err
	}

	return nil
}

package repository

import (
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}"
	"github.com/jinzhu/gorm"
)

type {{.CamelPkg}}Repository interface {
	FindAll() ([]*{{.Pkg}}.{{.CamelPkg}}, error)
	FindById(id int) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Store(*{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Update(int, *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error)
	Delete(id int) error
}

type {{.Pkg}}Repository struct {
	DB *gorm.DB
}

func New{{.CamelPkg}}Repository(DB *gorm.DB) {{.CamelPkg}}Repository {
	return &gorm{{.CamelPkg}}Repository{DB}
}

func (this *gorm{{.CamelPkg}}Repository) FindAll() ([]*{{.Pkg}}.{{.CamelPkg}}, error) {
	item := make([]*{{.Pkg}}.{{.CamelPkg}}, 0)
	if err := this.DB.Find(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (this *gorm{{.CamelPkg}}Repository) FindById(id int) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	item := &{{.Pkg}}.{{.CamelPkg}}{}
	if err := this.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (this *gorm{{.CamelPkg}}Repository) Store(value *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	if err := this.DB.Create(&value).Error; err != nil {
		return nil, err
	}
	return value, nil
}

func (this *gorm{{.CamelPkg}}Repository) Update(id int, value *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
        value.Id = id
	if err := this.DB.Update(&value).Error; err != nil {
		return nil, err
	}
	return value, nil
}

func (this *gorm{{.CamelPkg}}Repository) Delete(id int) error {
	item := &{{.Pkg}}.{{.CamelPkg}}{}
	if err := this.DB.Delete(&item, id).Error; err != nil {
		return err
	}

	return nil
}

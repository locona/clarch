package repository

import (
	"github.com/{{.CurrentUser}}/{{.CurrentRepo}}/{{.Pkg}}"
	"github.com/jinzhu/gorm"
)

type gorm{{.CamelPkg}}Repository struct {
	Conn *gorm.DB
}

func NewGorm{{.CamelPkg}}Repository(Conn *gorm.DB) {{.CamelPkg}}Repository {
	return &gorm{{.CamelPkg}}Repository{Conn}
}

func (this *gorm{{.CamelPkg}}Repository) FindAll() ([]*{{.Pkg}}.{{.CamelPkg}}, error) {
	item := make([]*{{.Pkg}}.{{.CamelPkg}}, 0)
	if err := this.Conn.Find(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (this *gorm{{.CamelPkg}}Repository) FindById(id int) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	var item {{.Pkg}}.{{.CamelPkg}}
	if err := this.Conn.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (this *gorm{{.CamelPkg}}Repository) Store(value *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	if err := this.Conn.Create(&value).Error; err != nil {
		return nil, err
	}
	return value, nil
}

func (this *gorm{{.CamelPkg}}Repository) Update(id int, value *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
        value.Id = id
	if err := this.Conn.Update(&value).Error; err != nil {
		return nil, err
	}
	return value, nil
}

func (this *gorm{{.CamelPkg}}Repository) Delete(id int) error {
	var item {{.Pkg}}.{{.CamelPkg}}
	if err := this.Conn.Delete(&item, id).Error; err != nil {
		return err
	}

	return nil
}

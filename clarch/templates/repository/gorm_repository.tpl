package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/{{.GithubUser}}/{{.GithubRepository}}/{{.Pkg}}"
)

type gorm{{.CamelPkg}}Repository struct {
	Conn *gorm.DB
}

func NewGorm{{.CamelPkg}}Repository(Conn *gorm.DB) {{.CamelPkg}}Repository {
	return &gorm{{.CamelPkg}}Repository{Conn}
}

func (this *gorm{{.CamelPkg}}Repository) FindAll() ([]*{{.Pkg}}.{{.CamelPkg}}, error) {
	o := make([]*{{.Pkg}}.{{.CamelPkg}}, 0)
	if err := this.Conn.Find(&o).Error; err != nil {
		return nil, err
	}
	return o, nil
}

func (this *gorm{{.CamelPkg}}Repository) FindById(id int) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	var o {{.Pkg}}.{{.CamelPkg}}
	if err := this.Conn.First(&o).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (this *gorm{{.CamelPkg}}Repository) Store(o *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	if err := this.Conn.Create(&o).Error; err != nil {
		return nil, err
	}
	return o, nil
}

func (this *gorm{{.CamelPkg}}Repository) Update(o *{{.Pkg}}.{{.CamelPkg}}) (*{{.Pkg}}.{{.CamelPkg}}, error) {
	if err := this.Conn.Update(&o).Error; err != nil {
		return nil, err
	}
	return o, nil
}

func (this *gorm{{.CamelPkg}}Repository) Delete(id int) error {
	var o {{.Pkg}}.{{.CamelPkg}}
	if err := this.Conn.Delete(&o, id).Error; err != nil {
		return err
	}

	return nil
}

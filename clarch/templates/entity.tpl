package {{.Pkg}}

import "time"

type {{.CamelPkg}} struct {
	Id int `json:"id" gorm:"primary_key"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at "sql:"index"`
}

package infra

import (
	"fmt"
	"log"

	config "github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Mysql *gorm.DB

func InitMysql() {
	host := config.Config.Mysql.Host
	port := config.Config.Mysql.Port
	user := config.Config.Mysql.User
	password := config.Config.Mysql.Password
	database := config.Config.Mysql.Database
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		user, password, host, port, database,
	))

	if err != nil {
		log.Fatal("MYSQL ERROR: ", err)
	}

	Mysql = db
}

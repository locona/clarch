package project

import (
	"log"

	"github.com/k0kubun/pp"
	"github.com/kelseyhightower/envconfig"
)

var Config *conf

type API struct {
	Port string `default:"8888"`
}

type Mysql struct {
	Host     string `required:"true" default:"127.0.0.1"`
	User     string `required:"true" default:"root"`
	Password string `required:"true" default:"password"`
	Database string `required:"true" default:"reckoner"`
	Port     string `default:"3306"`
}

type Redis struct {
	Host    string `required:"true" default:"127.0.0.1"`
	Port    string `required:"true" default:"6379"`
	MaxConn int    `required:"true" default:"1000"`
}

type conf struct {
	API
	Mysql
	Redis
}

func (this *conf) Print() {
	pp.Println(this)
}

func InitConfig() {
	api := API{}
	if err := envconfig.Process("api", &api); err != nil {
		log.Fatal(err)
	}

	mysql := Mysql{}
	if err := envconfig.Process("mysql", &mysql); err != nil {
		log.Fatal(err)
	}

	redis := Redis{}
	if err := envconfig.Process("redis", &redis); err != nil {
		log.Fatal(err)
	}

	Config = &conf{
		API:   api,
		Mysql: mysql,
		Redis: redis,
	}
}

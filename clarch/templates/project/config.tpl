package project

import (
	"log"

	"github.com/k0kubun/pp"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

type Logger struct{}

type conf struct {
	API
	Logger *zap.Config
	Mysql
}

func (this *conf) Print() {
	pp.Println(this)
}

func InitConfig() {
	var api API
	if err := envconfig.Process("api", &api); err != nil {
		log.Fatal(err)
	}

	var mysql Mysql
	if err := envconfig.Process("mysql", &mysql); err != nil {
		log.Fatal(err)
	}

	Config = &conf{
		API:   api,
		Mysql: mysql,
		Logger: func() *zap.Config {
			conf := zap.Config{
				Level:         zap.NewAtomicLevelAt(zap.DebugLevel),
				Development:   true,
				Encoding:      "console",
				EncoderConfig: zap.NewDevelopmentEncoderConfig(),
				Sampling: &zap.SamplingConfig{
					Initial:    100,
					Thereafter: 100,
				},
				OutputPaths: []string{
					"stdout",
				},
				ErrorOutputPaths: []string{
					"stderr",
				},
			}
			conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			return &conf
		}(),
	}
}

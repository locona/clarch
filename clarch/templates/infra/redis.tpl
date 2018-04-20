package infra

import (
	"log"

	config "github.com/{{.CurrentUser}}/{{.CurrentRepo}}/project"
	"github.com/garyburd/redigo/redis"
)

var Redis *redis.Pool

// redis のセッティング
func InitRedis() {
	host := config.Config.Redis.Host
	port := config.Config.Redis.Port
	maxConnections := config.Config.Redis.MaxConn
	conn := redis.NewPool(func() (redis.Conn, error) {
		conn, err := redis.Dial("tcp", host+":"+port)
		if err != nil {
			log.Fatal(err)
		}
		return conn, err
	}, maxConnections)

	Redis = conn
}

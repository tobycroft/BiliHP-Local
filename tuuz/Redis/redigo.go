package Redis

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"log"
)

var pool *redigo.Pool

func init() {
	redis_host := address
	redis_port := port
	pool_size := 10
	pool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%s", redis_host, redis_port))
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		return c, nil
	}, pool_size)
}

func Conn() redigo.Conn {
	return pool.Get()
}

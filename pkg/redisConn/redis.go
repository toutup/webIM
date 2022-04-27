package redisConn

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

const (
	// redis连接池最大空闲连接数
	maxIdle = 16
	// 连接池最大连接数，0表示无限制
	maxActive = 32
	// 空闲连接超时限制
	idleTimeout = 180 * time.Second
)

// 初始化redis
func Init() {
	fmt.Println(">开始初始化redis连接池...")
	if err := initRedis(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(">redis连接池初始化完成...")
}

func initRedis() error {
	if RedisConn == nil {
		RedisConn = &redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			IdleTimeout: idleTimeout,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")))
				if err != nil {
					return nil, err
				}
				if os.Getenv("REDIS_PASSWORD") != "" {
					if _, err := c.Do("AUTH", os.Getenv("REDIS_PASSWORD")); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return nil
}

// 获取onlineList连接 【database = 0】
func GetOnlineListConn() (redis.Conn, error) {
	conn := RedisConn.Get()
	_, err := conn.Do("select", 0)
	if err != nil {
		log.Fatal("err:", err.Error())
		return nil, err
	}
	return conn, nil
}

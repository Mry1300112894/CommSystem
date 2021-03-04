package model

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 声明一个连接池，Get方法 的接收者类型为 指针类型
func PoolInit(address string, maxIdle, maxActive int, idleTimeout time.Duration) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle: maxIdle, 		// 最大空闲连接数
		MaxActive: maxActive,		// 最大连接数，0表示没有限制
		IdleTimeout: idleTimeout,	// 最大空闲时间

		// 用于创建和配置连接的函数
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address)
		},
	}

	return pool
}

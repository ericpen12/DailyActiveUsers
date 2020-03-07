package main

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func getRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     "192.168.0.68:16379", // use default Addr
		Password: "password",           // no password set
		DB:       0,                    // use default DB
	}
}

func init() {
	rdb = redis.NewClient(getRedisConfig())
	if _, err := rdb.Ping().Result(); err != nil {
		panic("cannot connect redis: " + err.Error())
	}
}

func Sign(id int64) error {
	if err := rdb.SetBit(getActiveUserFormat(), id, 1).Err(); err != nil {
		return errors.New("cannot sign: " + err.Error())
	}
	return nil
}

func getActiveUserFormat() string {
	return "login:" + time.Now().Format("2006-01-02")
}

func GetActiveUserCount() (int64, error) {
	if count, err := rdb.BitCount(getActiveUserFormat(), nil).Result(); err != nil {
		return 0, errors.New("cannot get active user count: " + err.Error())
	} else {
		return count, nil
	}
}

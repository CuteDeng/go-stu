package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json unmashal err:", err)
		return
	}
	return
}

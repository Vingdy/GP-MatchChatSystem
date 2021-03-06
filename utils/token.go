package utils

import (
	"GP/model"
	"GP/redis"
	"encoding/json"
	"fmt"
)

func GetTokenInfo(token string) (userinfo model.User, err error) {
	jsoninfo, err := redis.Redis.Get(token).Bytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(jsoninfo, &userinfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

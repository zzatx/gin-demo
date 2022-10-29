package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	//不是用seed则每次生成的随机数都将一样
	rand.Seed(time.Now().Unix())
	for i := range result {
		//rand.Intn(n) 随机生成一个0到n-1的随机数 0 <= i < n
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

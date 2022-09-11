package login

import (
	"math/rand"
)

func GenOTP() string {

	bytes := make([]byte, 6)
	var pool = "1234567890abcdefghijklmopqrsyz"
	for i := 0; i < 6; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)

}

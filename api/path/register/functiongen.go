package register

import (
	db "api/db"

	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GanuserTag(s db.Db_mongo) string {
	for {
		bytes := make([]byte, 4)
		var pool = "1234567890"
		for i := 0; i < 4; i++ {
			bytes[i] = pool[rand.Intn(len(pool))]
		}
		some := make(chan primitive.D)
		//name + tag
		go s.Db_FindtOne("tag", string(bytes), some)
		ax := <-some
		if ax == nil || ax.Map()["identifind"] == false {
			return string(bytes)
		}
	}

}
func Ganuserid(s db.Db_mongo) string {
	for {
		var pool = "1234567890"
		dd := make([]byte, 13)
		for i := 0; i < 13; i++ {
			dd[i] = pool[rand.Intn(len(pool))]
		}
		some := make(chan primitive.D)
		go s.Db_FindtOne("userid", string(dd), some)
		ax := <-some

		if ax == nil || ax.Map()["identifind"] == false {
			return string(dd)
		}
	}

}
func GenOTP() string {

	bytes := make([]byte, 6)
	var pool = "1234567890abcdefghijklmopqrsyz"
	for i := 0; i < 6; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}
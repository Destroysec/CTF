package verify

import (
	db "api/db"
	//"fmt"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
)

func GenKEy() string {

	bytes := make([]byte, 19)
	var pool = "234567890abcdefghijxk2345678lmopqrsx2345678yzabxcdxefghxijkl2345678mopqrsyzx???>><<|\\}{!#@$%^&*())__"
	for i := 0; i < 19; i++ {

		bytes[i] = pool[rand.Intn(len(string(pool)))]
	}
	return string(bytes)

}
func SaveDAta(s db.Db_mongo, email, password, user, tag, time, id, jwt, tn string) {
	var post DATA
	session := make(map[string]string)
	session[tn] = jwt

	post.Time = time
	post.Email = email
	post.Subdata.Password = password
	post.Username = user
	post.UserId = id

	post.Tag = tag

	s.Db_InsertOneS(post)
	s.Db_FixOneStuck_UniDentify(bson.M{
		"email": bson.M{
			"$eq": email,
		},
	}, bson.M{"$push": bson.M{"sessionauthor": session}})

}

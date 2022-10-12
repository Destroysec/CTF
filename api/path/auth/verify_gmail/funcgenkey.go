package verify

import (
	db "api/all/db"
	//"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

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

	var f db.FormUD
	f.Db = &s
	f.Insert = post
	db.Db_insertOne(&f)
	s.Db_FixOneStuck_UniDentify(bson.M{
		"email": bson.M{
			"$eq": email,
		},
	}, bson.M{"$push": bson.M{"sessionauthor": session}})
	f = db.FormUD{}
	post = DATA{}
}

package register

import (
	db "api/db"

	r "api/random"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GanuserTag(s db.Db_mongo, un string) string {
	for {

		Tag := r.Randomid(4)
		some := make(chan primitive.D)
		name := make(chan primitive.D)
		//name + tag
		go s.Db_FindtOne("tag", Tag, some)
		go s.Db_FindtOne("username", un, name)
		ax := <-some
		cs := <-name
		if ax == nil || cs == nil {
			return Tag
		}
	}

}
func Ganuserid(s db.Db_mongo) string {
	for {

		id := r.Randomid(13)
		some := make(chan primitive.D)
		go s.Db_FindtOne("userid", string(id), some)
		ax := <-some

		if ax == nil || ax.Map()["identifind"] == false {
			return id
		}
	}

}

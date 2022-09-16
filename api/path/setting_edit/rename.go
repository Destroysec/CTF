package profile

import (
	db "api/db"
	d "api/jwt/service"

	h "api/hash_class"
	r "api/random"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type name struct {
	Password string `json:"password"`
	Newname  string `json:"newname"`
}

func GanuserTag(s db.Db_mongo, un string) string {
	for {

		Tag := r.Randomid(4)
		some := make(chan primitive.D, 100)
		name := make(chan primitive.D, 100)
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
func REname(c *gin.Context, s db.Db_mongo) {
	var ln name
	cha := make(chan primitive.D)
	dds := make(chan bool)
	if err := c.BindJSON(&ln); err != nil {
		c.AbortWithStatus(505)
		return
	}
	tokenHeader := c.Request.Header.Get("jwt")

	key, err := d.DecodeToken(tokenHeader)
	go s.Db_FindtOne("email", key.Claims.(jwt.MapClaims)["iss"].(string), cha)
	asd := (<-cha)
	go h.Vcheck(asd.Map()["subdata"].(primitive.D).Map()["password"].(string), ln.Password, dds)
	if err != nil {
		c.AbortWithStatus(505)
		return
	}
	asas := <-dds
	if !(asas) {
		c.JSON(405, gin.H{
			"message": asas,
		})

		return
	}
	tag := GanuserTag(s, asd.Map()["username"].(string))

	s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
		bson.M{"$set": bson.M{"username": ln.Newname}})
	go s.Db_FixOneStuck(bson.M{"username": ln.Newname, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
		bson.M{"$set": bson.M{"tag": tag}})
	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

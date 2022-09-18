package profile

import (
	db "api/db"
	d "api/jwt/service"

	h "api/hash_class"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pp struct {
	Subdata struct {
		Password string
	}
}
type password struct {
	Password       string `json:"password"`
	Newpassword    string `json:"newpassword"`
	Connewpassword string `json:"connewpassword"`
}

func ChangePassword(c *gin.Context, s db.Db_mongo) {
	var vea password
	cha := make(chan primitive.D)
	dds := make(chan bool)
	if err := c.BindJSON(&vea); err != nil {
		c.AbortWithStatus(503)
	}
	if vea.Newpassword != vea.Connewpassword {
		c.AbortWithStatus(500)
		return
	}
	if vea.Password == vea.Connewpassword || vea.Password == vea.Newpassword {
		c.AbortWithStatus(500)
		return
	}
	tokenHeader := c.Request.Header.Get("jwt")
	key, err := d.DecodeToken(tokenHeader)
	if err != nil {
		c.AbortWithStatus(504)
		return
	}
	go s.Db_FindtOne("email", key.Claims.(jwt.MapClaims)["iss"].(string), cha)
	asd := (<-cha)
	go h.Vcheck(asd.Map()["subdata"].(primitive.D).Map()["password"].(string), vea.Password, dds)
	if !(<-dds) {
		c.AbortWithStatus(501)
		return
	}
	s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
		bson.M{"$set": bson.M{"subdata.password": h.Mhash(vea.Newpassword)}})
	// key: asd.Map()["subdata"].(primitive.D).Map()["password"].(string)
}

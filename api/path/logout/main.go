package logout

import (
	db "api/db"
	d "api/jwt/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func Logout(c *gin.Context, s db.Db_mongo) {
	ds := c.Request.Header.Get("jwt")
	key, err := d.DecodeToken(ds)
	if err != nil {

		c.AbortWithStatus(505)
		return
	}
	fsd, err := s.Db_FindALLD("username", "email", key.Claims.(jwt.MapClaims)["iss"].(string), key.Claims.(jwt.MapClaims)["aud"].(string))
	if err != nil {

		c.AbortWithStatus(505)
		return
	}
	s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": fsd[0].Map()["username"]}, "tag": bson.M{"$eq": fsd[0].Map()["tag"]}, "email": bson.M{"$eq": fsd[0].Map()["email"]}},
		bson.M{"$unset": bson.M{"sessionauthor." + key.Claims.(jwt.MapClaims)["sub"].(string): ""}})
}

package Middleware

import (
	db "api/all/db"
	d "api/all/jwt/service"
	"fmt"

	h "api/all/hash_class"

	"github.com/golang-jwt/jwt"

	//"strings"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("jwt")
	if err := d.ValidateToken(s); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(505)
		return
	}
}
func MiddlewareSETTING(c *gin.Context, s db.Db_mongo) {
	ds := c.Request.Header.Get("jwt")
	adf := make(chan bool)
	key, err := d.DecodeToken(ds)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(505)
		return
	}
	fsd, err := s.Db_FindALLD("username", "email", key.Claims.(jwt.MapClaims)["iss"].(string), key.Claims.(jwt.MapClaims)["aud"].(string))
	if err != nil || fsd[0].Map()["sessionauthor"].(primitive.D).Map()[key.Claims.(jwt.MapClaims)["sub"].(string)] == nil {
		c.AbortWithStatus(404)
		return
	}
	go h.Vcheck(fsd[0].Map()["sessionauthor"].(primitive.D).Map()[key.Claims.(jwt.MapClaims)["sub"].(string)].(string), ds, adf)
	if !(<-adf) {
		c.AbortWithStatus(504)
		return
	}
	c.Next()

}

package profile

import (
	db "api/db"
	d "api/jwt/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func SETGithub(c *gin.Context, s db.Db_mongo) {
	file := c.PostForm("github")
	tokenHeader := c.Request.Header.Get("jwt")
	key, err := d.DecodeToken(tokenHeader)
	if err != nil {
		c.AbortWithStatus(505)
		return
	}
	s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
		bson.M{"$set": bson.M{"Github": file}})
	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

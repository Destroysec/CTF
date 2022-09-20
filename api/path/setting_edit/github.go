package profile

import (
	db "api/db"
	d "api/jwt/service"

	"net/http"
	"net/url"

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
	u, err := url.ParseRequestURI(file)
	if u.Hostname() == "github.com" {
		s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
			bson.M{"$set": bson.M{"subdata.Github": file}})
		c.JSON(http.StatusOK, gin.H{
			"message": "Your file has been successfully uploaded.",
		})
		return
	} else if u.Hostname() == "facebook.com" {
		s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
			bson.M{"$set": bson.M{"subdata.facebook": file}})
		c.JSON(http.StatusOK, gin.H{
			"message": "Your file has been successfully uploaded.",
		})
		return
	} else if u.Hostname() == "youtube.com" {
		s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
			bson.M{"$set": bson.M{"subdata.youtube": file}})
		c.JSON(http.StatusOK, gin.H{
			"message": "Your file has been successfully uploaded.",
		})
		return
	} else {
		c.AbortWithStatus(404)
		return
	}

	// File saved successfully. Return proper result

}

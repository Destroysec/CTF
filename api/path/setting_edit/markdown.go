package profile

import (
	db "api/db"
	d "api/jwt/service"
	r "api/random"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SETMarkdown(c *gin.Context, s db.Db_mongo) {
	file := c.PostForm("input")
	tokenHeader := c.Request.Header.Get("jwt")
	key, err := d.DecodeToken(tokenHeader)
	if err != nil {
		c.AbortWithStatus(505)
		return
	}
	if len(file) > 256 {
		c.AbortWithStatus(404)
		return
	}
	some := make(chan primitive.D, 100)
	ssss := make(chan primitive.D)
	a := r.Randomid(13)
	go s.Db_FindtOne("email", key.Claims.(jwt.MapClaims)["iss"].(string), ssss)
	dsf := (<-ssss)
	if dsf.Map()["Markdown"] != nil {
		os.Remove(".." + dsf.Map()["Markdown"].(string))
	}
	for {

		go s.Db_FindtOne("Markdown", "/Markdown/"+a+".md", some)
		if <-some == nil {
			break
		}
	}
	err = os.WriteFile("../Markdown/"+a+".md", []byte(file), 0644)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
		bson.M{"$set": bson.M{"Markdown": "/Markdown/" + a + ".md"}})
	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

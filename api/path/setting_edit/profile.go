package profile

import (
	db "api/db"
	d "api/jwt/service"
	r "api/random"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func SETProfile(c *gin.Context, s db.Db_mongo) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10*1024*1024)
	h, file, err := c.Request.FormFile("file")
	defer h.Close()
	tokenHeader := c.Request.Header.Get("jwt")
	key, err := d.DecodeToken(tokenHeader)
	if err != nil {
		c.AbortWithStatus(505)
		return
	}
	extension := filepath.Ext(file.Filename)
	things := []string{".jpeg", ".jpg", ".png", ".gif", ".JPEG", ".JPG", ".PNG", ".GIF"}
	if !contains(things, extension) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": ".jpeg .jpg .png .gif .JPEG .JPG .PNG .GIF",
		})
		return
	}
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	some := make(chan primitive.D, 100)
	ssss := make(chan primitive.D)
	a := r.Randomid(13)
	// The file is received, so let's save it

	go s.Db_FindtOne("email", key.Claims.(jwt.MapClaims)["iss"].(string), ssss)
	dsf := (<-ssss)
	if dsf.Map()["Profile"] != nil {
		os.Remove(".." + dsf.Map()["Profile"].(string))
	}
	for {

		go s.Db_FindtOne("Profile", "/IMG/"+a+extension, some)
		if <-some == nil {
			break
		}
	}
	if err := c.SaveUploadedFile(file, "../IMG/"+a+extension); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": key.Claims.(jwt.MapClaims)["aud"].(string)}, "tag": bson.M{"$eq": key.Claims.(jwt.MapClaims)["jti"].(string)}, "email": bson.M{"$eq": key.Claims.(jwt.MapClaims)["iss"].(string)}},
		bson.M{"$set": bson.M{"Profile": "/IMG/" + a + extension}})
	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

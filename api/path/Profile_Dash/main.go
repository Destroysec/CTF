package Dash

import (
	db "api/all/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Dash(c *gin.Context, s db.Db_mongo) {
	ID := c.Param("ID")
	ds, er := s.Db_FindALL("userid", ID)
	sil := make(map[string]string)
	if er != nil {
		c.AbortWithStatus(404)
		return
	}

	sil["username"] = ds[0]["username"].(string)
	sil["tag"] = ds[0]["tag"].(string)
	sil["userid"] = ds[0]["userid"].(string)

	if ds[0]["subdata"].(primitive.M)["profile"] != nil {
		sil["profile"] = ds[0]["subdata"].(primitive.M)["profile"].(string)
	}

	if ds[0]["subdata"].(primitive.M)["Github"] != nil {
		sil["Github"] = ds[0]["subdata"].(primitive.M)["Github"].(string)
	}
	if ds[0]["subdata"].(primitive.M)["facebook"] != nil {
		sil["facebook"] = ds[0]["subdata"].(primitive.M)["facebook"].(string)
	}

	if ds[0]["subdata"].(primitive.M)["youtube"] != nil {
		sil["youtube"] = ds[0]["subdata"].(primitive.M)["youtube"].(string)
	}

	if ds[0]["subdata"].(primitive.M)["markdown"] != nil {
		sil["markdown"] = ds[0]["subdata"].(primitive.M)["markdown"].(string)
	}

	c.JSON(200, gin.H{
		"message": sil,
	})
}

package profile

import (
	db "api/db"
	d "api/jwt/service"

	h "api/hash_class"
	// r "api/random"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	// "go.mongodb.org/mongo-driver/bson"
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
		c.AbortWithStatus(505)
	}
	tokenHeader := c.Request.Header.Get("jwt")
	key, err := d.DecodeToken(tokenHeader)
	if err == nil {
		c.AbortWithStatus(505)
	}
	go s.Db_FindtOne("email", key.Claims.(jwt.MapClaims)["iss"].(string), cha)
	asd := (<-cha)
	go h.Vcheck(asd.Map()["subdata"].(primitive.D).Map()["password"].(string), vea.Password, dds)
	if !(<-dds) {
		c.AbortWithStatus(505)
	} else if vea.Newpassword != vea.Connewpassword {
		c.AbortWithStatus(505)
	}

}

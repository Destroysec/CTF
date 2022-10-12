package profile

import (
	"api/all/db"
	"api/all/gmail"
	h "api/all/hash_class"
	d "api/all/jwt/service"
	rr "api/all/random"
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type t struct {
	Email    string
	Username string
	Tag      string
	password string
}
type urls struct {
	Username string `json:"name" url:"username"`
}
type from struct {
	Nawpassword string `json:"newpassword"`
	Repassword  string `json:"repassword"`
}

func (T *t) GenURL(c *gin.Context, s db.Db_mongo, am gmail.GAmll) {
	re := rr.Randomid(25)
	go s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": T.Username}, "tag": bson.M{"$eq": T.Tag}, "email": bson.M{"$eq": T.Email}},
		bson.M{"$set": bson.M{"sessionp.Callback": re}})
	go s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": T.Username}, "tag": bson.M{"$eq": T.Tag}, "email": bson.M{"$eq": T.Email}},
		bson.M{"$set": bson.M{"sessionp.newpassword": h.Mhash(T.password)}})
	go s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": T.Username}, "tag": bson.M{"$eq": T.Tag}, "email": bson.M{"$eq": T.Email}},
		bson.M{"$set": bson.M{"sessionp.TImeout": time.Now().Format("2006-01-02 15:04:05")}})
	go am.SEndCAll(T.Username, T.Tag, "http://"+c.Request.Host+"/callBack/"+re, T.Email)
}

func SendURL(c *gin.Context, s db.Db_mongo, am gmail.GAmll) {
	tokenHeader := c.Request.Header.Get("jwt")
	var body from
	d, err := d.DecodeToken(tokenHeader)
	if err != nil {
		c.AbortWithStatus(504)
		return
	}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(404, err)
		return
	}
	if len(body.Nawpassword) < 8 {
		c.JSON(404, gin.H{
			"message": "asdsss",
		})
		return
	}
	if body.Nawpassword != body.Nawpassword {
		c.JSON(404, gin.H{
			"message": "asdsss",
		})
		return
	}
	B := t{
		Username: d.Claims.(jwt.MapClaims)["aud"].(string),
		Email:    d.Claims.(jwt.MapClaims)["iss"].(string),
		Tag:      d.Claims.(jwt.MapClaims)["jti"].(string),
		password: body.Nawpassword,
	}
	go B.GenURL(c, s, am)
	c.JSON(200, gin.H{
		"message": "OKay",
	})
}
func CheckURL(c *gin.Context, s db.Db_mongo) {

	item, err := s.Db_FindALL("sessionp.Callback", c.Param("username"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	t, _ := time.Parse("2006-01-02 15:04:05", item[0]["sessionp"].(primitive.M)["TImeout"].(string))
	fmt.Println("")
	if t.Before(t.Add(time.Second * time.Duration(60))) {
		s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": item[0]["username"]}, "tag": bson.M{"$eq": item[0]["tag"]}, "email": bson.M{"$eq": item[0]["email"]}},
			bson.M{"$set": bson.M{"subdata.password": item[0]["sessionp"].(primitive.M)["newpassword"]}})
		s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": item[0]["username"]}, "tag": bson.M{"$eq": item[0]["tag"]}, "email": bson.M{"$eq": item[0]["email"]}},
			bson.M{"$unset": bson.M{"sessionp": ""}})
		c.JSON(200, gin.H{
			"message": "OKay",
		})
	} else {
		c.JSON(305, gin.H{
			"message": "sd",
		})
		s.Db_FixOneStuck(bson.M{"username": bson.M{"$eq": item[0]["username"]}, "tag": bson.M{"$eq": item[0]["tag"]}, "email": bson.M{"$eq": item[0]["email"]}},
			bson.M{"$unset": bson.M{"sessionp": ""}})
		return
	}

}

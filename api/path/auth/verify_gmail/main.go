package verify

import (
	db "api/all/db"

	h "api/all/hash_class"

	"fmt"

	j "api/all/jwt/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
   type StandardClaims struct {
       Audience  string `json:"aud,omitempty"`
       ExpiresAt int64  `json:"exp,omitempty"`
       Id        string `json:"jti,omitempty"`
       IssuedAt  int64  `json:"iat,omitempty"`
       Issuer    string `json:"iss,omitempty"`
       NotBefore int64  `json:"nbf,omitempty"`
       Subject   string `json:"sub,omitempty"`
   }
*/
//reg

func Verifyotp(c *gin.Context, s db.Db_mongo) {
	var Ax GEtheader
	checkotp := make(chan bool)
	if err := c.BindJSON(&Ax); err != nil {
		fmt.Println(err)
	}
	a, err := j.DecodeToken(Ax.Jwt)
	if err != nil {
		fmt.Println(err)
	}

	hashjwt := make(chan string)

	tn := time.Now().Format("2006_01_02-15:04:05")
	some, e := s.Db_FindALLunD("username", "tag", a.Claims.(jwt.MapClaims)["jti"].(string), a.Claims.(jwt.MapClaims)["aud"].(string))

	if e != nil {

		some, e = s.Db_FindALLD("username", "tag", a.Claims.(jwt.MapClaims)["jti"].(string), a.Claims.(jwt.MapClaims)["aud"].(string))

		if e != nil || some[0].Map()["SessionOTP"] == nil {
			c.JSON(404, gin.H{
				"message": "fill s",
			})
		}
		go h.Vcheck(some[0].Map()["SessionOTP"].(string), Ax.Jwt+" "+Ax.OTP, checkotp)
		g, _ := j.GenerateTokenV(c, some[0].Map()["tag"].(string), some[0].Map()["username"].(string), tn, some[0].Map()["email"].(string), int64(222))
		go h.AsyncMhash(g, hashjwt)
		fddf := <-checkotp
		if fddf {
			asdsa := <-hashjwt
			go s.Db_FixOneStuck(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}},
				bson.M{"$set": bson.M{"sessionauthor." + tn: asdsa}})

			c.JSON(200, gin.H{
				"message": "req okay",
				"s":       g,
			})
		} else {
			c.JSON(401, gin.H{
				"message": "fild",
			})
		}

	} else {

		if some[0].Map()["count"].(int32) == 3 {
			go s.Db_Delete_UniDentify(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}})
		}
		go h.Vcheck(some[0].Map()["sessionreg"].(string), Ax.Jwt+" "+Ax.OTP, checkotp)

		g, _ := j.GenerateTokenV(c, some[0].Map()["tag"].(string), some[0].Map()["username"].(string), tn, some[0].Map()["email"].(string), int64(222))
		go h.AsyncMhash(g, hashjwt)
		if <-checkotp {
			asdsa := <-hashjwt
			SaveDAta(s, some[0].Map()["email"].(string), some[0].Map()["subdata"].(primitive.D).Map()["password"].(string), some[0].Map()["username"].(string), some[0].Map()["tag"].(string), some[0].Map()["time"].(string), some[0].Map()["userid"].(string), g, tn)
			go s.Db_Delete_UniDentify(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}})
			go s.Db_FixOneStuck(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}},
				bson.M{"$set": bson.M{"sessionauthor." + tn: asdsa}})
			c.JSON(200, gin.H{
				"message": "req okay",
				"s":       g,
			})

		} else {
			s.Db_FixOneStuck_UniDentify(bson.M{
				"email": bson.M{
					"$eq": some[0].Map()["email"].(string),
				},
			}, bson.M{"$inc": bson.M{"count": 1}})
			c.JSON(401, gin.H{
				"message": "fild",
			})
		}
	}

}

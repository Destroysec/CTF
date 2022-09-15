package verify

import (
	db "api/db"

	h "api/hash_class"

	"fmt"
	"strings"

	j "api/jwt/service"
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
	var split []string

	tn := time.Now().Format("2006-01-02 15:04:05")
	some, e := s.Db_FindALLunD("username", "tag", a.Claims.(jwt.MapClaims)["jti"].(string), a.Claims.(jwt.MapClaims)["aud"].(string))

	if e != nil {
		some, _ = s.Db_FindALLD("username", "tag", a.Claims.(jwt.MapClaims)["jti"].(string), a.Claims.(jwt.MapClaims)["aud"].(string))

		split = strings.Split(some[0].Map()["SessionOTP"].(string), " ")
		go h.Vcheck(split[1], Ax.OTP, checkotp)

		g, _ := j.GenerateTokenReg(c, some[0].Map()["tag"].(string), some[0].Map()["username"].(string), some[0].Map()["email"].(string), tn)
		go h.AsyncMhash(g, hashjwt)
		fddf := <-checkotp
		if fddf {
			session := make(map[string]string)
			asdsa := <-hashjwt
			session[tn] = asdsa
			go s.Db_FixOneStuck(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}},
				bson.M{"$push": bson.M{"sessionauthor": session}})
			s.RemoveArray(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}},
				bson.M{"$set": bson.M{"SessionOTP": some[0].Map()["SessionOTP"].(string)}})
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
		split := strings.Split(some[0].Map()["sessionreg"].(string), " ")
		if some[0].Map()["count"].(int32) == 3 {
			go s.Db_Delete_UniDentify(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}})
		}
		go h.Vcheck(split[1], Ax.OTP, checkotp)

		g, _ := j.GenerateTokenReg(c, some[0].Map()["tag"].(string), some[0].Map()["username"].(string), some[0].Map()["email"].(string), tn)
		go h.AsyncMhash(g, hashjwt)
		if <-checkotp {
			session := make(map[string]string)
			go SaveDAta(s, some[0].Map()["email"].(string), some[0].Map()["subdata"].(primitive.D).Map()["password"].(string), some[0].Map()["username"].(string), some[0].Map()["tag"].(string), some[0].Map()["time"].(string), some[0].Map()["userid"].(string), g, tn)

			go s.Db_Delete_UniDentify(bson.M{"email": bson.M{"$eq": some[0].Map()["email"].(string)}})
			session[tn] = <-hashjwt
			s.Db_FixOneStuck_UniDentify(bson.M{
				"email": bson.M{
					"$eq": some[0].Map()["email"].(string),
				},
			}, bson.M{"$push": bson.M{"sessionauthor": session}})
			c.JSON(200, gin.H{
				"message": "req okay",
				"s":       g,
			})

		} else {
			s.Db_FixOneStuck_UniDentify(bson.M{
				"email": bson.M{
					"$eq": some[0].Map()["email"].(string),
				},
			}, bson.M{"$set": bson.M{"count": some[0].Map()["count"].(int32) + 1}})
			c.JSON(401, gin.H{
				"message": "fild",
			})
		}
	}

}

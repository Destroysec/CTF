package login

import (
	db "api/db"
	"api/gmail"
	h "api/hash_class"
	jwt "api/jwt/service"

	r "api/random"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DATA struct {
	Sessions_DATA struct {
		session string
	}
}

func Login(c *gin.Context, s db.Db_mongo, am gmail.GAmll) {
	var fromreg ln
	t := time.Now().Format("2006-01-02 15:04:05")
	OTP := r.RandomOTP(6)
	cha := make(chan primitive.D)
	jwthash := make(chan string)
	Sessionhash := make(chan string)
	ds := make(chan bool)
	if err := c.BindJSON(&fromreg); err != nil {
		c.JSON(404, gin.H{
			"message": "fail"})
		return

	}
	go s.Db_FindtOne("email", fromreg.Email, cha)
	password := fromreg.Password
	key := <-cha
	if key != nil {

		go h.Vcheck(key.Map()["subdata"].(primitive.D).Map()["password"].(string), password, ds)
		if <-ds {
			go h.AsyncMhash(OTP, Sessionhash)
			g, _ := jwt.GenerateToken(c, key.Map()["tag"].(string), key.Map()["username"].(string), string(t), key.Map()["email"].(string), int64(46))
			go h.AsyncMhash(g, jwthash)
			go s.Db_FixOneStuck(bson.M{"email": bson.M{"$eq": key.Map()["email"].(string)}, "username": bson.M{"$eq": key.Map()["username"].(string)}}, bson.M{"$set": bson.M{"SessionOTP": <-jwthash + " " + <-Sessionhash}})
			go am.SEndlogin(key.Map()["username"].(string), key.Map()["tag"].(string), OTP, key.Map()["email"].(string))

			c.JSON(200, gin.H{
				"message": "login suss",
				"jwt":     g,
			})
		} else {

			c.JSON(404, gin.H{
				"message": "login fail"})
		}
	} else {

		c.JSON(404, gin.H{
			"message": "email not font"})
	}

}

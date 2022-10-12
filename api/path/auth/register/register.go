package register

import (
	db "api/all/db"
	"api/all/gmail"
	h "api/all/hash_class"
	jwt "api/all/jwt/service"
	r "api/all/random"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reg struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}
type DATA struct {
	Email                       string
	Username, Tag, UserId, Time string
	SessionReg                  string
	Identifind                  bool
	Count                       int
	Subdata                     struct {
		Password string
	}
}

func saveDAta(s db.Db_mongo, email, password, user, tag, SEss, time string) {
	var post DATA
	post.Time = time
	post.Email = email
	post.Subdata.Password = h.Mhash(password)
	post.Username = user
	post.UserId = string(Ganuserid(s))
	post.SessionReg = SEss
	post.Tag = tag
	post.Identifind = false
	post.Count = 0
	go s.Db_InsertOneS_UniDentify(post)
	post = DATA{}
}

func Register(c *gin.Context, s db.Db_mongo, am gmail.GAmll) {
	var fromreg Reg
	if err := c.BindJSON(&fromreg); err != nil {
		c.JSON(401, gin.H{
			"message": "err",
		})
	}

	cha := make(chan primitive.D)

	hs := make(chan string)
	hss := make(chan string)
	go s.Db_FindtOne("email", fromreg.Email, cha)

	if fromreg.Password == fromreg.Repassword {
		datauser := <-cha
		if datauser == nil || datauser.Map()["identifind"] == false {
			go s.Db_DeleteMany_UniDentify(bson.M{"email": bson.M{"$eq": fromreg.Email}})
			tag := GanuserTag(s, fromreg.Username)
			Ax := r.RandomOTP(6)
			go h.AsyncMhash(Ax, hss)
			t := time.Now().Format("2006-01-02 15:04:05")
			g, _ := jwt.GenerateTokenReg(c, tag, fromreg.Username, t, fromreg.Email, int64(47))
			go h.AsyncMhash(g, hs)
			go saveDAta(s, fromreg.Email, fromreg.Password, fromreg.Username, tag, <-hs+" "+<-hss, t)
			go am.SEndlogin(fromreg.Username, tag, Ax, fromreg.Email)

			fromreg = Reg{}
			fmt.Println(Ax)
			c.JSON(200, gin.H{
				"message": "Register suss",
				"hee":     g,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Register fg",
			})
		}

	} else {
		c.JSON(402, gin.H{
			"message": "password!=repassword fail",
		})
	}

}

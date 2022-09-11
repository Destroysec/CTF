package register

import (
	db "api/db"
	"api/gmail"
	h "api/hash_class"
	jwt "api/jwt/service"

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
			tag := string(GanuserTag(s))
			Ax := GenOTP()
			go h.AsyncMhash(Ax, hss)
			t := time.Now().Format("2006-01-02 15:04:05")
			g, _ := jwt.GenerateTokenReg(c, tag, fromreg.Username, fromreg.Email, t, int64(60456))
			go h.AsyncMhash(g, hs)
			go saveDAta(s, fromreg.Email, fromreg.Password, fromreg.Username, tag, <-hs+" "+<-hss, t)
			go am.SEndlogin(fromreg.Username, tag, Ax, fromreg.Email)
			c.JSON(200, gin.H{
				"message": "Register suss",
				"hee":     g,
			})
			fromreg = Reg{}
			fmt.Println(Ax)
		}

	} else {
		c.JSON(402, gin.H{
			"message": "password!=repassword fail",
		})
	}

}

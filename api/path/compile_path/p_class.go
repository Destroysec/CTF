package path

import (
	db "api/db"
	"api/gmail"
	p "api/path/login"
	P "api/path/register"

	v "api/path/verify_gmail"
	config "api/setting"

	//"fmt"
	"github.com/gin-gonic/gin"
)

var s db.Db_mongo
var am gmail.GAmll

func init() {

	so := config.Get_Config()
	ztructDB := db.DBStarterConfig{DATA: so}
	s.Db_start(ztructDB)
	am.LoginCon(so)
}
func M(c *gin.Context) {

	Leve := c.PostForm("leve")
	Asm := c.PostForm("asm")

	if Leve == "1" {
		if Asm == "hackerman" {
			c.JSON(200, gin.H{
				"message": Asm + "leve:" + Leve + " True",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "leve:" + Leve + " False",
			})
		}
	}

}
func Verifyotp_func(c *gin.Context) {
	v.Verifyotp(c, s)
}

func Register(c *gin.Context) {

	P.Register(c, s, am)
}
func Login(c *gin.Context) {

	p.Login(c, s, am)
}

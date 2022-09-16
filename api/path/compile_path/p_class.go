package path

import (
	db "api/db"
	"api/gmail"
	p "api/path/login"
	Pa "api/path/register"
	SET "api/path/setting_edit"
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

	Pa.Register(c, s, am)
}
func Login(c *gin.Context) {

	p.Login(c, s, am)
}
func SETProfile(c *gin.Context) {

	SET.SETProfile(c, s)
}
func SETMarkdown(c *gin.Context) {

	SET.SETMarkdown(c, s)
}
func SETGithub(c *gin.Context) {

	SET.SETGithub(c, s)
}
func Rename(c *gin.Context) {

	SET.REname(c, s)
}

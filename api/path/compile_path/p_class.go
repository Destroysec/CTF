package path

import (
	"api/Middleware"
	db "api/all/db"
	"api/all/gmail"
	config "api/all/setting"
	pay "api/path/Pay"
	Dash "api/path/Profile_Dash"
	p "api/path/auth/login"
	"api/path/auth/logout"
	Pa "api/path/auth/register"
	v "api/path/auth/verify_gmail"
	SET "api/path/setting_edit"

	//"fmt"
	"github.com/gin-gonic/gin"
)

var s db.Db_mongo
var am gmail.GAmll
var asd *config.Data_Config

func init() {

	asd = config.Get_Config()

	ztructDB := db.DBStarterConfig{DATA: asd}
	s.Db_start(ztructDB)
	am.LoginCon(asd)
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

func Logout(c *gin.Context) {

	logout.Logout(c, s)
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
func ChangePassword(c *gin.Context) {

	SET.ChangePassword(c, s)
}
func Sendreset(c *gin.Context) {

	SET.SendURL(c, s, am)
}
func Callback(c *gin.Context) {

	SET.CheckURL(c, s)
}
func Dassh(c *gin.Context) {
	Dash.Dash(c, s)

}
func PAy(c *gin.Context) {
	pay.Getgv(c, s, asd.Pay.Numberphone, asd.Pay.Reql)

}

// middleware
func MS(c *gin.Context) {
	Middleware.MiddlewareSETTING(c, s)
}

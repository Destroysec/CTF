package main

import (
	//jwt "api/jwt/service"
	M "api/Middleware"
	p "api/path/compile_path"
	rr "api/random"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/julianshen/gin-limiter"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,X-API-KEY, jwt,otp")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
func main() {
	r := gin.Default()
	var a string
	r.Use(CORSMiddleware())
	go func() {
		for {

			a = rr.RandomKEY(25)
			os.WriteFile("../.env", []byte("KEY="+a), 0644)
			fmt.Println(a)
			time.Sleep(time.Minute * 10)
		}
	}()
	lm := limiter.NewRateLimiter(time.Minute, 10, func(ctx *gin.Context) (string, error) {
		key := ctx.Request.Header.Get("X-API-KEY")
		if key == a {
			return key, nil
		}
		return "", errors.New("API key is missing")
	})
	api := r.Group("/", M.AuthorizationMiddleware)
	apilogin := r.Group("/apilogin")
	apilogin.POST("/ln", lm.Middleware(), p.Login) //fix this
	apilogin.POST("/reg", lm.Middleware(), p.Register)
	api.POST("/q", lm.Middleware(), p.M)
	api.POST("/verifyotp", lm.Middleware(), p.Verifyotp_func)
	api.POST("/setProfile", p.SETProfile)
	//apilogin.GET("/Check", p.C)
	r.Run(":9000")

	// listen and serve on 0.0.0.0:9000 (for windows "localhost:9000")
	/*
		#Example
		import requets
		requets.post("localhost:9000/q",data={"leve":"1","ams":"hackerman"})

	*/

}

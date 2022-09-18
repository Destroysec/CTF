package Middleware

import (
	db "api/db"
	jwt "api/jwt/service"
	"fmt"

	//"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("jwt")
	if err := jwt.ValidateToken(s); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(505)
		return
	}
}
func MiddlewareSETTING(c *gin.Context, s db.Db_mongo) {
	ds := c.Request.Header.Get("jwt")
	if err := jwt.ValidateToken(ds); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(505)
		return
	}
}

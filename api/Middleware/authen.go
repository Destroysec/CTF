package Middleware

import (
	jwt "api/jwt/service"
	"fmt"

	//"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("jwt")

	//token := strings.TrimPrefix(s, "Bearer ")
	//_ , err := jwt.DecodeToken(token)

	if err := jwt.ValidateToken(s); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(505)
		return
	}
	c.Next()
}

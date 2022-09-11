package Middleware

import (
	jwt "api/jwt/service"
	"fmt"

	//"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(c *gin.Context) {
	var request DataJwT
	err := c.ShouldBindHeader(&request)
	if err != nil {
		fmt.Print(err)
		return
	}
	s := c.Request.Header.Get("jwt")

	//token := strings.TrimPrefix(s, "Bearer ")
	//_ , err := jwt.DecodeToken(token)

	if err := jwt.ValidateToken(s); err != nil {
		fmt.Print(err)
		c.AbortWithStatus(505)
		return
	}
}

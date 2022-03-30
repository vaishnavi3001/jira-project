package handlers

import (
	"fmt"
	ct "jira-backend/constants"
	dt "jira-backend/dbutils"
	ut "jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthInterceptor(c *gin.Context) {
	tokenCookie, err := c.Request.Cookie(ct.Access_token)
	if err != nil {
		fmt.Println("Here")
		c.AbortWithStatusJSON(http.StatusUnauthorized, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
		return
	}

	tokenStr := tokenCookie.Value

	claims, err := ut.ParseJwtToken(tokenStr)
	fmt.Println(claims)
	if err != nil {
		if err == jwt.ErrSignatureInvalid || err == ct.ErrTokenInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
			return
		}
	}

	if !dt.CheckTokenInDb(tokenStr, claims.UserId) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
		return
	}
	c.Set(ct.USER_ID, claims.UserId)
	c.Next()
}

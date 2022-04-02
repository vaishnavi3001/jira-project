package handlers

import (
	"fmt"
	ct "jira-backend/constants"
	dt "jira-backend/dbutils"
	ut "jira-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthInterceptor(c *gin.Context) {
	tokenStr := ""
	tokenCookie, err := c.Request.Cookie(ct.Access_token)

	if err != nil {
		headerToken := c.Request.Header.Get("Authorization")
		if len(headerToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
			return
		}
		tokenArr := strings.Split(headerToken, " ")
		tokenStr = tokenArr[1]
		fmt.Println(tokenStr)
		if len(tokenStr) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
			return
		}
	} else {
		tokenStr = tokenCookie.Value
	}

	claims, err := ut.ParseJwtToken(tokenStr)

	if err != nil {
		if err == jwt.ErrSignatureInvalid || err == ct.ErrTokenInvalid {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
			return
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, ut.GetErrorResponse(ct.EXPIRED_TOKEN))
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

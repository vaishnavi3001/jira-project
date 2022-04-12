package handlers

import (
	"fmt"
	ct "jira-backend/constants"
	dt "jira-backend/dbutils"
	sk "jira-backend/skeletons"
	ut "jira-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Userlogin(c *gin.Context) {
	var req sk.UserAuthReq

	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	uid := dt.CheckLoginCreds(req.Username, req.Password)
	if uid == -1 {
		c.JSON(http.StatusOK, ut.GetErrorResponse(ct.INVALID_CREDENTIALS))
		return
	}

	tokenStr, tokenErr := ut.CreateJwtToken(uint(uid))

	if tokenErr != nil {
		c.JSON(http.StatusOK, ut.GetErrorResponse(ct.LOGIN_FAILURE))
	}

	dt.AddUserToken(tokenStr, uint(uid))
	c.SetCookie(ct.Access_token, tokenStr, 216000, "/", ".jira-clone.com", false, true)
	c.JSON(http.StatusOK, ut.GetSuccessResponse(ct.LOGIN_SUCCESSFUL, sk.LoginSuccessResp{Access_token: tokenStr}))
}

func UserLogout(c *gin.Context) {
	tokenCookie, err := c.Request.Cookie(ct.Access_token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ut.GetErrorResponse(ct.ACTION_NOT_AUTHORIZED))
		return
	}

	tokenStr := tokenCookie.Value

	dt.DeleteToken(tokenStr)
	c.SetCookie(ct.Access_token, "", 0, "/", ".jira-clone.com", false, true)
	c.JSON(http.StatusOK, ut.GetSuccessResponse(ct.LOGOUT_SUCCESSFUL, ""))
}

func UserRegister(c *gin.Context) {
	var req sk.UserRegister

	if err := c.BindJSON(&req); err != nil {
		ut.ThrowBadRequest(c)
		return
	}

	c.JSON(http.StatusOK, dt.RegisterUser(req))
}

func GetUserProfile(c *gin.Context) {
	user_id := ut.GetUserIdFromContext(c)
	fmt.Println(user_id)
	c.JSON(http.StatusOK, dt.GetProfileUser(user_id))
}
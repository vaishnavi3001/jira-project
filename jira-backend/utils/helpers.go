package utils

import (
	"fmt"
	ct "jira-backend/constants"
	sk "jira-backend/skeletons"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetSuccessResponse(msg string, val interface{}) gin.H {
	return gin.H{
		ct.Status:   true,
		ct.Message:  msg,
		ct.Response: val,
	}
}

func GetErrorResponse(error_msg string) gin.H {
	return gin.H{
		ct.Status:  false,
		ct.Message: error_msg,
	}
}

func ThrowBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, GetErrorResponse(ct.BAD_REQUEST))
	c.AbortWithStatus(http.StatusBadRequest)
}

func GetResponse(status bool, msg string, val interface{}) gin.H {
	return gin.H{
		ct.Status:   status,
		ct.Message:  msg,
		ct.Response: val,
	}
}

func DateFormat(dateStr string) (time.Time, error) {
	t, err := time.Parse(ct.SPRINT_DATE_FORMAT, dateStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t, err
}

func CreateJwtToken(userId uint) (string, error) {
	secretKey := []byte(Vconfig.GetString("creds.secret_key"))
	duration := Vconfig.GetInt64("creds.duration")
	currTime := time.Now()

	claims := sk.Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currTime.Unix(),
			ExpiresAt: currTime.Add(time.Duration(duration) * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func ParseJwtToken(tokenSt string) (sk.Claims, error) {
	claims := &sk.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenSt, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Vconfig.GetString("creds.secret_key")), nil
	})

	if err != nil {
		return *claims, err
	}

	if !tkn.Valid {
		err = ct.ErrTokenInvalid
	}

	return *claims, err
}

func GetUserIdFromContext(c *gin.Context) uint {
	return c.GetUint(ct.USER_ID)
}

func EncryptEmailInvite(emailId string, projectID uint) (string, error) {
	secretKey := []byte(Vconfig.GetString("mail_config.mail_encryption_key"))
	expiryDuration := Vconfig.GetInt64("mail_config.duration") * 24
	currTime := time.Now()
	claims := sk.InviteClaim{
		EmailId:   emailId,
		ProjectId: projectID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currTime.Unix(),
			ExpiresAt: currTime.Add(time.Duration(expiryDuration) * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func ParseEmailInvite(encryptedLink string) (sk.InviteClaim, error) {
	claims := &sk.InviteClaim{}

	tkn, err := jwt.ParseWithClaims(encryptedLink, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Vconfig.GetString("mail_config.mail_encryption_key")), nil
	})

	if err != nil {
		return *claims, err
	}

	if !tkn.Valid {
		err = ct.ErrLinkExpired
	}

	return *claims, err
}

func InviteDateToString(timestamp time.Time) string {
	return timestamp.Format("2006-01-02 15:04:05.000")
}

func InviteDateStringToDate(dateStr string) (time.Time, error) {

	t, err := time.Parse("2006-01-02 15:04:05.000", dateStr)
	return t, err
}

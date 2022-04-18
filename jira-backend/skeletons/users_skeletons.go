package skeletons

import "github.com/golang-jwt/jwt"

type UsersBaseReq struct {
	UserId uint `json:"user_id"`
}

type UserAuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserIdReq struct {
	UserId int `json:"user_id"`
}

type UserEntry struct {
	UserId   uint `json:"user_id"`
	UserName uint `json:"username"`
	UserRole uint `json:"user_role"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterReq struct {
	Firstname string
}

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

type UserRegister struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email_id"`
}

type LoginSuccessResp struct {
	Access_token string `json:"access_token"`
}

type UserProfile struct {
	UserId    uint   `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email_id"`
}

type ChangePasswordReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
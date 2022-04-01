package tests

import (
	ct "jira-backend/constants"
	"net/http"
	"time"
)

type testBody struct {
	init         func(*http.Request)
	url          string
	method       string
	bodyData     string
	expectedCode int
	expectedResp string
}

//{"message":"","resp":{"project_name":"Project 5","project_id":5,"created_at":"2022-03-31T19:12:09.045909-04:00"},"status":true}
const TokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODAzMTI1MTQsImlhdCI6MTY0ODc3NjUxNH0.kWl7lZpDsywQDUGKeEG4h7d_P9nF4HvoUnNUhzQXVsk"

func setTokenRequestInCookie(req *http.Request) {
	cookie := http.Cookie{Name: ct.Access_token, Value: TokenStr, Expires: time.Now().Add(365 * 24 * time.Hour), HttpOnly: true}
	req.AddCookie(&cookie)
}

var testData = []testBody{
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/create",
		"POST",
		`{"name":"Project 1"}`,
		http.StatusOK,
		`{"message":"","resp":{"project_name":"Project 1","project_id":([0-9]+),"created_at":"([A-Z0-9:\-\.]+)"},"status":true}`,
	},
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/create",
		"POST",
		``,
		http.StatusBadRequest,
		`{"message":"BAD_REQUEST","status":false}`,
	},
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/list",
		"POST",
		``,
		http.StatusOK,
		`{"message":"","resp":{"projects":[{"name":"Project 1","id":[0-9]+,"created_at":"([A-Z0-9:\-\.]+)","user_role":[0-9]+}]},"status":true}`,
	},
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/info",
		"POST",
		``,
		http.StatusBadRequest,
		`{"message":"BAD_REQUEST","status":false}`,
	},
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/info",
		"POST",
		`{"project_id": 1}`,
		http.StatusOK,
		`{"message":"","resp":{"project_id":([0-9]+),"project_name":"Project 1","owner_uname":"pypalkar23","owner_id":([0-9]+),"owner_fname":"Mandar","owner_lname":"Palkar","created_at":"([A-Z0-9:\-\.]+)"},"status":true}`,
	},
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/members",
		"POST",
		`{"project_id": 1}`,
		http.StatusOK,
		`{"message":"","resp":{"members":[{"user_id":[0-9]+,"first_name":"Mandar","last_name":"Palkar","user_role":[0-9]+}]},"status":true}`,
	},
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/members",
		"POST",
		``,
		http.StatusBadRequest,
		`{"message":"BAD_REQUEST","status":false}`,
	},
}

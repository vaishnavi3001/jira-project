package tests

import (
	ct "jira-backend/constants"
	"jira-backend/dbutils"
	"jira-backend/models"
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
	afterfunc    func() bool
}

const (
	TokenStr         = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODAzMTI1MTQsImlhdCI6MTY0ODc3NjUxNH0.kWl7lZpDsywQDUGKeEG4h7d_P9nF4HvoUnNUhzQXVsk"
	badRequestString = `{"message":"BAD_REQUEST","status":false}`
)

func setTokenRequestInCookie(req *http.Request) {
	cookie := http.Cookie{Name: ct.Access_token, Value: TokenStr, Expires: time.Now().Add(365 * 24 * time.Hour), HttpOnly: true}
	req.AddCookie(&cookie)
}

var apiTestData = []testBody{
	//Project Create Test - Success
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/create",
		"POST",
		`{"name":"Project 1"}`,
		http.StatusOK,
		`{"message":"","resp":{"project_name":"Project 1","project_id":([0-9]+),"created_at":"([A-Z0-9:\-\.]+)"},"status":true}`,
		func() bool {
			var count int64
			dbutils.DB.Where("role_id = 1 AND user_id = 1 AND project_id = 1").Find(&models.UserRole{}).Count(&count)
			return count == 1
		},
	},
	//Project Create Test Bad Request
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/create",
		"POST",
		``,
		http.StatusBadRequest,
		badRequestString,
		func() bool { return true },
	},
	//Project List API Correct Request-Response
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/list",
		"POST",
		``,
		http.StatusOK,
		`{"message":"","resp":{"projects":[{"name":"Project 1","id":[0-9]+,"created_at":"([TZ0-9:\-\.]+)","user_role":[0-9]+}]},"status":true}`,
		func() bool { return true },
	},
	//Project List API Bad Request
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/info",
		"POST",
		``,
		http.StatusBadRequest,
		badRequestString,
		func() bool { return true },
	},
	//Project List API Correct Request Response
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/info",
		"POST",
		`{"project_id": 1}`,
		http.StatusOK,
		`{"message":"","resp":{"project_id":([0-9]+),"project_name":"Project 1","owner_uname":"pypalkar23","owner_id":([0-9]+),"owner_fname":"Mandar","owner_lname":"Palkar","created_at":"([A-Z0-9:\-\.]+)"},"status":true}`,
		func() bool { return true },
	},
	//List Project Members - Correct Request Response
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/members",
		"POST",
		`{"project_id": 1}`,
		http.StatusOK,
		`{"message":"","resp":{"members":[{"user_id":[0-9]+,"first_name":"Mandar","last_name":"Palkar","user_role":[0-9]+}]},"status":true}`,
		func() bool { return true },
	},
	//Project Members - Bad Request
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/members",
		"POST",
		``,
		http.StatusBadRequest,
		badRequestString,
		func() bool { return true },
	},
	//Project Members - Bad Request
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/members",
		"POST",
		``,
		http.StatusBadRequest,
		badRequestString,
		func() bool { return true },
	},

	/*sprint*/
	/*sprint create bad request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/create",
		"POST",
		``,
		http.StatusBadRequest,
		badRequestString,
		func() bool { return true },
	},
	/*sprint create success request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/create",
		"POST",
		`{"sprint_name" :"Sprint 1","project_id" : 1,"start_date": "2021-03-07","end_date" : "2021-03-21"}`,
		http.StatusOK,
		`{"message":"","resp":{"sprint_name":"Sprint 1","sprint_id":[0-9]+},"status":true}`,
		func() bool { return true },
	},
	/*sprint create with same name*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/create",
		"POST",
		`{"sprint_name" :"Sprint 1","project_id" : 1,"start_date": "2021-03-07","end_date" : "2021-03-21"}`,
		http.StatusOK,
		`{"message":"SPRINT_EXISTS","status":false}`,
		func() bool { return true },
	},
	/*creating another sprint*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/create",
		"POST",
		`{"sprint_name" :"Sprint 2","project_id" : 1,"start_date": "2021-03-07","end_date" : "2021-03-21"}`,
		http.StatusOK,
		`{"message":"","resp":{"sprint_name":"Sprint 2","sprint_id":[0-9]+},"status":true}`,
		func() bool { return true },
	},
	/*sprint list correct*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/list",
		"POST",
		`{"project_id":1}`,
		http.StatusOK,
		`{"message":"","resp":{"sprints":[{"name":"Sprint 1","id":[0-9]+,"start_date":"[TZ0-9:\-\.]+","end_date":"[TZ0-9:\-\.]+","created_at":"[TZ0-9:\-\.]+","project_id":[0-9]+},{"name":"Sprint 2","id":[0-9]+,"start_date":"[TZ0-9:\-\.]+","end_date":"[TZ0-9:\-\.]+","created_at":"[TZ0-9:\-\.]+","project_id":[0-9]+}]},"status":true}`,
		func() bool { return true },
	},
	/*sprint delete*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/delete",
		"POST",
		`{"sprint_id":2}`,
		http.StatusOK,
		`{"message":"SPRINT_DELETE_SUCCESS","resp":"","status":true}`,
		func() bool { return true },
	},
	/*sprint delete again*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/delete",
		"POST",
		`{"sprint_id":2}`,
		http.StatusOK,
		`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
		func() bool { return true },
	},
	/*sprint list after deletion*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/list",
		"POST",
		`{"project_id":1}`,
		http.StatusOK,
		`{"message":"","resp":{"sprints":[{"name":"Sprint 1","id":1,"start_date":"[TZ0-9:\-\.]+","end_date":"[TZ0-9:\-\.]+","created_at":"[TZ0-9:\-\.]+","project_id":1}]},"status":true}`,
		func() bool { return true },
	},
	/*sprint info*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/info",
		"POST",
		`{"sprint_id":1}`,
		http.StatusOK,
		`{"message":"","resp":{"sprint_id":1,"sprint_name":"Sprint 1","created_at":"[TZ0-9:\-\.]+","start_date":"[TZ0-9:\-\.]+","end_date":"[TZ0-9:\-\.]+"},"status":true}`,
		func() bool { return true },
	},
	/*sprint update success*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/sprint/update",
		"POST",
		`{"sprint_name" :"Sprint Name Changed","sprint_id": 1,"start_date": "2021-03-06","end_date" : "2021-03-21"}`,
		http.StatusOK,
		`{"message":"SPRINT_UPDATE_SUCCESS","resp":"","status":true}`,
		func() bool { return true },
	},
	//sprint update without token
	{func(req *http.Request) {
	},
		"/api/sprint/update",
		"POST",
		`{"sprint_name" :"Sprint Name Changed","sprint_id": 1,"start_date": "2021-03-06","end_date" : "2021-03-21"}`,
		http.StatusUnauthorized,
		`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
		func() bool { return true },
	},
	/*****Sprint Test Cases End*****/
	/*****Issues test cases start ****/
	/*issue creation without appropriate token*/
	{func(req *http.Request) {
	},
		"/api/issue/create",
		"POST",
		`{"issue_title":"issue title 1","issue_description": "sample issue description","issue_type": 1,"creator": 2,"assignee": 1,"sprint_id": 1,"project_id": 1}`,
		http.StatusUnauthorized,
		`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
		func() bool { return true },
	},
	/*issue create correct request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/create",
		"POST",
		`{"issue_title":"issue title 1","issue_description":"sample issue description","issue_type": 1,"creator": 2,"assignee": 1,"sprint_id": 1,"project_id": 1}`,
		http.StatusOK,
		`{"message":"","resp":{"issue_id":0,"issue_title":"issue title 1","issue_description":"sample issue description","issue_type":1,"creator_name":"pypalkar23","assignee_name":"pypalkar23","created_at":"[TZ0-9:\-\.]+","issue_status":1},"status":true}`,
		func() bool { return true },
	},
	/*issue create bad request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/create",
		"POST",
		`{}`,
		http.StatusOK,
		`{"message":"INVALID_USER","status":false}`,
		func() bool { return true },
	},
	/*issue create bad request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/create",
		"POST",
		`{}`,
		http.StatusOK,
		`{"message":"INVALID_USER","status":false}`,
		func() bool { return true },
	},
	/*issue list bad request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/list",
		"POST",
		``,
		http.StatusBadRequest,
		badRequestString,
		func() bool { return true },
	},
	/* issue list with invalid sprint id*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/list",
		"POST",
		`{}`,
		http.StatusOK,
		`{"message":"INVALID_SPRINT","status":false}`,
		func() bool { return true },
	},
	/* issue list correct request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/list",
		"POST",
		`{"sprint_id":1}`,
		http.StatusOK,
		`{"message":"","resp":{"issues":[{"issue_id":1,"title":"issue title 1","status":1,"created_at":"[TZ0-9:\-\.]+"}]},"status":true}`,
		func() bool { return true },
	},
	/*issue info correct request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/info",
		"POST",
		`{"issue_id":2}`,
		http.StatusOK,
		`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
		func() bool { return true },
	},
	/*issue info correct request*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/issue/info",
		"POST",
		`{"issue_id":1}`,
		http.StatusOK,
		`{"message":"","resp":{"issue_id":[0-9]+,"title":"issue title 1","type":[0-9]+,"sprint_id":[0-9]+,"sprint_name":"Sprint Name Changed","project_id":[0-9]+,"description":"sample issue description","assignee_id":[0-9]+,"assignee_name":"pypalkar23","creator_id":[0-9]+,"creator_name":"pypalkar23","created_at":"[TZ0-9:\-\.]+","project_name":"Project 1","issue_status":[0-9]+},"status":true}`,
		func() bool { return true },
	},
	/*****Issues Test Cases End*****/
	/*****User Profile test cases start ****/
	/*View user profile without appropriate token*/
	{func(req *http.Request) {
	},
		"/api/user/info",
		"GET",
		``,
		http.StatusUnauthorized,
		`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
		func() bool { return true },
	},
	/*View user profile with appropriate token*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/user/info",
		"GET",
		``,
		http.StatusOK,
		`{"message":"USER_FOUND","resp":{"user_id":0,"username":"pypalkar23","firstname":"Mandar","lastname":"Palkar","email_id":"mandar.palkar23@gmail.com"},"status":true}`,
		func() bool { return true },
	},
	// /*Set username to a username that already exists*/
	// {func(req *http.Request) {
	// 	setTokenRequestInCookie(req)
	// },
	// 	"/api/user/info",
	// 	"PATCH",
	// 	`{"username":"amhaske32","firstname":"Mandar","lastname":"Palkar","email_id": "mandar.palkar23@gmail.com"}`,
	// 	http.StatusOK,
	// 	`{"message":"USER_FOUND","resp":{"user_id":0,"username":"pypalkar23","firstname":"Mandar","lastname":"Palkar","email_id":"mandar.palkar23@gmail.com"},"status":true}`,
	// 	func() bool { return true },
	// },

	/*Set username to a username that does not exist*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/user/info",
		"PATCH",
		`{"username":"palkar","firstname":"Mandar","lastname":"Palkar","email_id": "mandar.palkar23@gmail.com"}`,
		http.StatusOK,
		`{"message":"USER_PROFILE_UPDATE_SUCCESS","resp":{"user_id":0,"username":"palkar","firstname":"Mandar","lastname":"Palkar","email_id":"mandar.palkar23@gmail.com"},"status":true}`,
		func() bool { return true },
	},

	/*Set username to a username that does not exist*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/user/info",
		"PATCH",
		`{"username":"pypalkar23","firstname":"Mandar","lastname":"Palkar","email_id": "mandar.palkar23@gmail.com"}`,
		http.StatusOK,
		`{"message":"USER_PROFILE_UPDATE_SUCCESS","resp":{"user_id":0,"username":"pypalkar23","firstname":"Mandar","lastname":"Palkar","email_id":"mandar.palkar23@gmail.com"},"status":true}`,
		func() bool { return true },
	},

	// /*Set email to a email that already exists in db*/
	// {func(req *http.Request) {
	// },
	// 	"/api/issue/create",
	// 	"POST",
	// 	`{"issue_title":"issue title 1","issue_description": "sample issue description","issue_type": 1,"creator": 2,"assignee": 1,"sprint_id": 1,"project_id": 1}`,
	// 	http.StatusUnauthorized,
	// 	`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
	// 	func() bool { return true },
	// },

	/*Set email to an email that does not exist in db*/
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/user/info",
		"PATCH",
		`{"username":"pypalkar23","firstname":"Mandar","lastname":"Palkar","email_id": "mandar.palkar13@gmail.com"}`,
		http.StatusOK,
		`{"message":"USER_PROFILE_UPDATE_SUCCESS","resp":{"user_id":0,"username":"pypalkar23","firstname":"Mandar","lastname":"Palkar","email_id":"mandar.palkar13@gmail.com"},"status":true}`,
		func() bool { return true },
	},

	

}

var loginTestData = []testBody{
	//User Login Request Incorrect Credentials
	{func(req *http.Request) {
	},
		"/login",
		"POST",
		`{"username":"pypalkar23","password":"wrong password"}`,
		http.StatusOK,
		`{"message":"INVALID_CREDENTIALS","status":false}`,
		func() bool { return true },
	},
	//User Logout Request
	{func(req *http.Request) {
	},
		"/login",
		"POST",
		`{"username":"pypalkar23","password":"dd29b8cb089a56606fca480e137c27c4"}`,
		http.StatusOK,
		`{"message":"LOGIN_SUCCESSFUL","resp":{"access_token":"[A-Za-z0-9\._\-]+"},"status":true}`,
		func() bool { return true },
	},
	//Successful Logout
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/logout",
		"GET",
		``,
		http.StatusOK,
		`{"message":"LOGOUT_SUCCESSFUL","resp":"","status":true}`,
		func() bool { return true },
	},
	//Invalid Logout
	{func(req *http.Request) {

	},
		"/logout",
		"GET",
		``,
		http.StatusUnauthorized,
		`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
		func() bool { return true },
	},
}

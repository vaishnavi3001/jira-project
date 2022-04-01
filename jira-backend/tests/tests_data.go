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
const (
	TokenStr         = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODAzMTI1MTQsImlhdCI6MTY0ODc3NjUxNH0.kWl7lZpDsywQDUGKeEG4h7d_P9nF4HvoUnNUhzQXVsk"
	badRequestString = `{"message":"BAD_REQUEST","status":false}`
)

func setTokenRequestInCookie(req *http.Request) {
	cookie := http.Cookie{Name: ct.Access_token, Value: TokenStr, Expires: time.Now().Add(365 * 24 * time.Hour), HttpOnly: true}
	req.AddCookie(&cookie)
}

var testData = []testBody{
	//Project Create Test - Success
	{func(req *http.Request) {
		setTokenRequestInCookie(req)
	},
		"/api/project/create",
		"POST",
		`{"name":"Project 1"}`,
		http.StatusOK,
		`{"message":"","resp":{"project_name":"Project 1","project_id":([0-9]+),"created_at":"([A-Z0-9:\-\.]+)"},"status":true}`,
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
	},
	//sprint update without token
	{func(req *http.Request) {
	},
		"/api/sprint/update",
		"POST",
		`{"sprint_name" :"Sprint Name Changed","sprint_id": 1,"start_date": "2021-03-06","end_date" : "2021-03-21"}`,
		http.StatusUnauthorized,
		`{"message":"ACTION_NOT_AUTHORIZED","status":false}`,
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
	},
}

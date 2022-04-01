package tests

import (
	"bytes"
	"fmt"
	"jira-backend/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	InitConfig()
	DBFree()
	DBInitForTest()
	DBInsertRows()
	exitval := m.Run()
	DBFree()
	os.Exit(exitval)
}

func TestApis(t *testing.T) {
	fmt.Printf("--- Running suite of %d test cases for API testing ---\n", len(apiTestData))
	r := routes.SetupRouter(true)
	for _, testentry := range apiTestData {
		asserts := assert.New(t)
		bodyData := testentry.bodyData
		req, err := http.NewRequest(testentry.method, testentry.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		testentry.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testentry.expectedCode, w.Code)
		asserts.Regexp(testentry.expectedResp, w.Body.String())
		asserts.Equal(true, testentry.afterfunc())
	}
}

func TestLoginLogout(t *testing.T) {
	r := routes.SetupRouter(true)
	fmt.Printf("--- Running suite of %d test cases for Login/Logout ---\n", len(loginTestData))
	for _, testentry := range loginTestData {
		asserts := assert.New(t)
		bodyData := testentry.bodyData
		req, err := http.NewRequest(testentry.method, testentry.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		testentry.init(req)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testentry.expectedCode, w.Code)
		asserts.Regexp(testentry.expectedResp, w.Body.String())
		asserts.Equal(true, testentry.afterfunc())
	}
}

package tests

import (
	"bytes"
	"jira-backend/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackend(t *testing.T) {

	r := routes.SetupRouter(true)
	for _, testentry := range testData {
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
	}
}

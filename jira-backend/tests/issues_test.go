package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListTest(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	jsonparam := `{"user_id":1,"issue_id":6}`

	_, err := http.NewRequest("GET", "/issues/info", strings.NewReader(string(jsonparam)))

	assert.Nil(t, err)
	assert.Equal(t, 200, w.Code, "OK")
}

func TestIssueUpdate(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	jsonparam := `{
		"user_id":1,
		"issue_id":5,
		"sprint_id":1,
		"issue_name":"Sample Issue 7",
		"issue_text":"Thing abcdfcer is not working",
		"creator":1,
		"assignee":2,
		"project_id":1
	}`

	_, err := http.NewRequest("POST", "/issues/update", strings.NewReader(string(jsonparam)))

	assert.Nil(t, err)
	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code, "OK")
}

func TestIssueList(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	jsonparam := `{
		" "project_id":1,
		"sprint_id":1
	}`

	_, err := http.NewRequest("POST", "/issues/list", strings.NewReader(string(jsonparam)))

	assert.Nil(t, err)
	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code, "OK")
}

func TestIssueView(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	jsonparam := `{
		"user_id":1,
		"issue_id":5,
	}`

	_, err := http.NewRequest("POST", "/issues/delete", strings.NewReader(string(jsonparam)))

	assert.Nil(t, err)
	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code, "OK")
}

func TestProjectView(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	jsonparam := `{
		"user_id":1,
		"project_id":5
	}`

	_, err := http.NewRequest("POST", "/project/info", strings.NewReader(string(jsonparam)))

	assert.Nil(t, err)
	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code, "OK")
}

func TestProjectCreate(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	jsonparam :=
		`{
			"name" :"Boring Company",
			"user_id":2
		}`

	_, err := http.NewRequest("POST", "/projects/create", strings.NewReader(string(jsonparam)))

	assert.Nil(t, err)
	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code, "OK")
}

func TestProjectDelete(t *testing.T) {

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	jsonparam :=
		`{
			"project_id":4,
			"user_id":2
		}`

	_, err := http.NewRequest("POST", "/projects/delete", strings.NewReader(string(jsonparam)))

	assert.Nil(t, err)
	t.Log(w.Body.String())
	assert.Equal(t, 200, w.Code, "OK")
}

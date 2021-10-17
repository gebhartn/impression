package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gebhartn/impress/utils"
	"github.com/stretchr/testify/assert"
)

func TestSignUpSuccess(t *testing.T) {
	tearDown()
	setup()

	var req = `{"user":{"username":"nicholas","password":"pass"}}`

	r := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(req))

	r.Header.Set("Content-Type", "application/json")

	h.Register(e)

	res, _ := e.Test(r, -1)
	if assert.Equal(t, http.StatusCreated, res.StatusCode) {
		body, _ := ioutil.ReadAll(res.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "nicholas", m["username"])
		assert.NotEmpty(t, m["token"])
	}
}

func TestLoginSuccess(t *testing.T) {
	tearDown()
	setup()

	var req = `{"user":{"username":"testuser","password":"testpass"}}`

	r := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(req))
	r.Header.Set("Content-type", "application/json")

	h.Register(e)

	res, _ := e.Test(r, -1)
	if assert.Equal(t, http.StatusOK, res.StatusCode) {
		body, _ := ioutil.ReadAll(res.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "testuser", m["username"])
	}
}

func TestCurrentUserCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	r := httptest.NewRequest(http.MethodGet, "/api/user", nil)
	r.Header.Set("Content-type", "application/json")
	r.Header.Set("Authorization", authHeader(utils.GenerateJWT(1)))
	h.Register(e)
	res, err := e.Test(r, -1)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, res.StatusCode) {
		body, _ := ioutil.ReadAll(res.Body)
		m := responseMap(body, "user")
		assert.Equal(t, "testuser", m["username"])
		assert.NotEmpty(t, m["token"])
	}
}

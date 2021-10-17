package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

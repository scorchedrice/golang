package myapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()

	// GET method test
	req, _ := http.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Index", string(data))
}

func TestIndexPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()

	// GET method test
	req, _ := http.NewRequest("GET", "/bar", nil)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Worlds", string(data))
	//	Fail!
}

func TestIndexPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()

	// GET method test
	req, _ := http.NewRequest("GET", "/bar?name=scorchedrice", nil)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello scorchedrices", string(data))
	//	Fail!
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()

	// GET method test
	req, _ := http.NewRequest("POST", "/foo", strings.NewReader(`{"first_name":"scorchedrice", "last_name":"master", "email":"asd@gmail.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("scorchedrice", user.FirstName)
	assert.Equal("master", user.LastName)

}

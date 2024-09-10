package myapp

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello world", string(data))
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "No Users")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "No User id:89")

	resp, err = http.Get(ts.URL + "/users/56")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User id:56")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"Han", "last_name":"jiwoong", "email":"woongveloper98@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(int(id)))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// Delete는 RESTFUL API 이후 등장한 개념으로 기본으로 제공되진 않음.
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "No User id:1")

	// post하고 delete해보는 과정
	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"Han", "last_name":"jiwoong", "email":"woongveloper98@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = io.ReadAll(resp.Body)
	assert.Contains(string(data), "Deleted User id:1")
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id":1, "first_name":"updated", "last_name":"updated", "email":"updated@gmail.com"}`))
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	fmt.Println(data)
	assert.Contains(string(data), "No User id:1")

	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"Han", "last_name":"jiwoong", "email":"woongveloper98@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"updated"}`, user.ID)

	req, _ = http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(updateStr))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(user.ID, updateUser.ID)
	assert.Equal("updated", updateUser.FirstName)
	assert.Equal(user.LastName, updateUser.LastName)
	assert.Equal(user.Email, updateUser.Email)
}

func TestUser_WithUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"Han", "last_name":"jiwoong", "email":"woongveloper98@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"Two", "last_name":"jiwoong", "email":"woongveloper98@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	users := []*User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))
}

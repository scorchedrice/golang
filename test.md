# Test 환경
## http Handler 묶기
- main.go의 내용을 다른 .go폴더에 저장
```go
// myapp/go
package myapp

import (
  "encoding/json"
  "fmt"
  "net/http"
  "time"
)

type User struct {
  FirstName string    `json:"first_name"`
  LastName  string    `json:"last_name"`
  Email     string    `json:"email"`
  CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  user := new(User)
  err := json.NewDecoder(r.Body).Decode(user)
  if err != nil {
    // Error 이 발생한경우
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, err)
    return
  }
  user.CreatedAt = time.Now()
  data, _ := json.Marshal(user)
  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Query().Get("name")
  if name == "" {
    name = "World"
  }
  fmt.Fprintf(w, "Hello %s", name)

}

func NewHttpHandler() http.Handler {
  mux := http.NewServeMux()
  // mux :
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world")
  })

  mux.HandleFunc("/bar", barHandler)

  mux.Handle("/foo", &fooHandler{})
  return mux
}
```
- 이를 main.go에서 다음과 같이 사용
```go
package main

import (
  "golang/myapp"
  "net/http"
)

func main() {
  http.ListenAndServe(":8080", myapp.NewHttpHandler())
}
```

## go에서의 testcode 작성
- test code : app_test.go
    - _뒤에 test를 붙혀서 컨벤션 지키기
- smartystreets/goconvey, testify와 같은 유용한 도구들 존재
    - go get github/... 해당 github에 설치법 따라 설치 진행

### Test 예시 (Get, '/')
```go
package myapp

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
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
```

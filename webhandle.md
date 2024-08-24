# 간단한 웹 서버
```go
package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Bar")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	http.HandleFunc("/bar", barHandler)

	http.Handle("/foo", &fooHandler{})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
```
- 8080 포트, Hello world 출력
- Fprint(w, "~~") : writer 에다가 print
- 각 주소별 차이 비교해보기

# mux : 왜 쓰는가
```go
func main() {
	mux := http.NewServeMux()
	// mux :
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
```
- /api/v1... 처럼 주소를 커스텀 한다고 생각하자 (api용 web용 따로 둘 수 있잖아)
- 독립적인 라우터를 만들 수 있다.

# query 활용
```go
func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)

}
```
- url : http://localhost:8080/bar?name=scorchedrice
    - Hello scorchedrice

# api호출 관련 기초내용
```go
package main

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

func main() {
	mux := http.NewServeMux()
	// mux :
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
```
- json 타입일 때 변수명을 go에게 알려줘야 잘 알아들음
    - go와 json의 컨벤션 문제
    - w.Header에 해당 타입이 무엇인지 알려줘야 return이 json임을 인지함
    - 요청을 보낼 때 어떤식으로 처리하는지 fooHandler 분석해볼 필요성 있음.

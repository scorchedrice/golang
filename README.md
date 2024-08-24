# Go를 학습하는 이유
- 새로운 언어를 학습하고 싶었다.
  - 상대적으로 학습하기 쉽다는 평이 많아 시도해볼만 하다고 생각했다.
  - 실제로 docker, 당근마켓, 쿠버네티스, 발로란트 등 많은 어플리케이션에서 사용한다.
- 웹서버 및 Microservice, 클라우드 서비스 등에 활용 가능하다.
  - FrontEnd이지만 관심이 있는 언어를 학습하며 관련 내용을 익힐 수 있을 것이라 생각했다.

# Hello Go
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello Go")
}
```

# 변수선언
```go
var a int = 20
```
- var, const 는 dart, javascript와 큰 차이 없음.
- 변수명 뒤에 type을 지정
- 물론 a := 20 으로하면 알아서 유추해서 할 수 있음.

# 타입 변환 (string => int)
- strconv를 import해서 변환할 수 있음.
```go
// a가 string으로 '5'인경우

b := strconv.Atoi(a)
// b := strconv.parseInt(a,10,64)
// 반대로 int를 str으로 하려면 Itoa
```


# 배열
```go
array := [3]int{1,2,3}
```
- 배열의 길이를 정하고 ([3])
- 뒤에 배열 지정
- 물론 []으로 하고 하면 유동적인 길이를 가질 수 있다.

# Map
```go
map := make(map[string]int)
map["age"] = 90
```

# 포인터
```go
a := 10
b := &a

fmt.Println(b)
```
- 10이란 값이 a라는 변수명으로 저장되어 있는데, b는 그 메모리 주소를 할당 받는다.

# go 명령어
- 동시실행을 가능하게 하는 명령어

# terminal 명령어
- go run test.go
- go build test.go : exe파일로 빌드

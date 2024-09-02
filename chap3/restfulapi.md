# ioutil
- 사용 자제할 것
  - Go 자체에서 사용을 자제하라고 함
    - 메모리 누수 등의 이슈
  - io에 ReadAll과 같은 동일 기능 존재하니 그것 사용할 것

# {바뀌는 값}을 가지는 경우 test
- id, pk ...
  - 강의에서는 gorilla mux활용

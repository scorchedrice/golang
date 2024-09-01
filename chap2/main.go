package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	// w, http.ResponseWriter 는 클라이언트로 응답을 보낼 때 사용
	// r, http.Request 는 요청 정보를 담고있음.

	// index.html의 upload_file name입니다.
	uploadfile, header, err := r.FormFile("upload_file")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	// 함수가 종료되기 전, 파일을 닫도록 보장, 리소스 누수 방지
	defer uploadfile.Close()

	// 저장할 디렉토리 생성, 0777은 모든 사용자에게 읽기 쓰기 실행권한을 준다는 의미
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)

	// 업로드한 파일의 전체 경로를 생성하고
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	//새 파일을 생성하고, 생성 실패인 경우 500 반환
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// 업로드한 파일 내용을 새로 생성한 파일에 복사합니다.
	io.Copy(file, uploadfile)

	// 성공적인 경우 200 반환
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	// uploads로 오는 요청은 uploadsHandler로 처리합니다.
	http.HandleFunc("/uploads", uploadsHandler)
	// "/"으로 요청이 오면 ./static/ 을 서빙합니다.
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":3030", nil)
}

package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)

	// 테스트할 파일의 경로 지정, 파일을 열고 테스트가 끝나면 파일을 닫습니다.
	path := "/Users/hanjiwoong/Desktop/golang/chap2/uploads/한지웅_취업사진.JPG"
	file, _ := os.Open(path)
	defer file.Close()

	// 테스트하는 환경을 초기화하는 과정으로, 기존 업로드 디렉토리를 삭제합니다.
	os.RemoveAll("./uploads")

	// multipart 형식의 HTTP 요청 본문을 생성합니다. 파일 업로드 요청을 시뮬레이션 합니다.
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	// 테스트용 HTTP 요청과 응답 객체를 생성합니다.
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	uploadsHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)

	// upload된 파일이 실제로 있는지 확인합니다.
	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	// upload된 파일이 원본 파일과 동일한지 확인하는 과정입니다.
	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(uploadData, originData)
}

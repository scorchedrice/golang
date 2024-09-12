package main

import (
	"fmt"
	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var receivedData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	// send data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	// 암호화 된 component (EncrptComponent)
	self.com.Operator(string(encryptData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData))
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	receivedData = data
}

func main() {
	// 서로 묶이고 묶인 관계
	sender := &EncryptComponent{key: "abcde", com: &ZipComponent{com: &SendComponent{}}}
	sender.Operator("Hello world")
	fmt.Println(
		"압축된 Hello world",
		sentData,
	)

	receiver := &UnzipComponent{com: &DecryptComponent{key: "abcde", com: &ReadComponent{}}}
	receiver.Operator(sentData)
	fmt.Println("받은 Hello world : ", receivedData)
}

package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Person struct {
	Name string
	Age int
}

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			fmt.Printf("Done err is: %v\n", ctx.Err())
			return ctx.Err()
		}
	}
}

type demo struct {

}

func (demo) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (demo) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func TestDefer()(int, error){
	defer func() {
		fmt.Println("defer")
	}()
	return fmt.Println("return")
}
type PubgCompetition struct {
	username string `json:"playerName"`
	matchmode int8 `json:"matchModel"`
	matchid string `json:"id"`
	playat string `json:"playedAt"`
}
var cookies []*http.Cookie

// AESCFBdecrypt AES CFB 解密
func AESCFBdecrypt1(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	s := strconv.Itoa(0xA0)
	fmt.Printf("%T, %v\n", s, s)

}

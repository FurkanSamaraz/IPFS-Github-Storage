package block

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	path = "./pullipfs/"
)
var (
	pathdown = "./downipfs/"
)

func Blockencrypt(path string, key string) {

	dizin, err := os.Open(path)
	if err != nil {
		fmt.Println("Dizin bulunamadı!")
		os.Exit(1)
	}
	defer dizin.Close()
	liste, _ := dizin.Readdirnames(0) // Açıklamada okuyun
	for _, isim := range liste {
		dosya, _ := ioutil.ReadFile(path + isim)

		a, _ := Encrypt(dosya, key)
		veris := []byte(a)
		ioutil.WriteFile(path+isim, veris, 0644)

	}
}
func Blockdecryption(path string, key string) {
	dizin, err := os.Open(pathdown)
	if err != nil {
		fmt.Println("Dizin bulunamadı!")
		os.Exit(1)
	}
	defer dizin.Close()
	liste, _ := dizin.Readdirnames(0) // Açıklamada okuyun
	for _, isim := range liste {
		dosya, _ := ioutil.ReadFile(pathdown + isim)

		decText, _ := Decrypt(string(dosya), key)
		veri := []byte(decText)
		ioutil.WriteFile(pathdown+isim, veri, 0644)

	}
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text []byte, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := text
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

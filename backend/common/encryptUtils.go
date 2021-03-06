package common

import (
	"crypto/aes"
	"encoding/base64"
	"io"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"crypto/md5"
	"golang.org/x/crypto/scrypt"
	"github.com/SSSaaS/sssa-golang"
	"fmt"
)

func Encrypt(password, secret, userId string) ([]byte, error) {
	passwordHash, _ := scrypt.Key([]byte(password), []byte(userId), 16384, 8, 8, 32)
	block, err := aes.NewCipher(passwordHash)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString([]byte(secret))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(password string, secret []byte, userId string) ([]byte, error) {
	passwordHash := GetScryptHash(password, userId)
	block, err := aes.NewCipher(passwordHash)
	if err != nil {
		return nil, err
	}
	if len(secret) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := secret[:aes.BlockSize]
	secret = secret[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(secret, secret)
	data, err := base64.StdEncoding.DecodeString(string(secret))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SplitSecret(secret string){
	shares, _ := sssa.Create(2,3,secret)
	// Securely send shares to their owners
	// For prototype just print shares to console
	for _, i := range shares {
		fmt.Println(i)
	}
}

func GetMd5Hash(text string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hasher.Sum(nil)
}

func GetScryptHash(password, salt string) []byte {
	passwordHash, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 8, 32)
	return passwordHash
}

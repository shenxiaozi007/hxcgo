package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Test_JWT(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["token"] = "test"

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["fakeExpiresAt"] = time.Now().Add(time.Second * 15).Unix() //update token per 15min

	tokenStr, err := token.SignedString([]byte("mysign"))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(tokenStr)

}

func Test_DecodeJWT(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjcwODExNTIsImZha2VFeHBpcmVzQXQiOjE1NjY4MjE5NjcsInRva2VuIjoidGVzdCJ9.PxSollpoq8VzPUXzyajW71y7liSeu1uQ_yTRToZZn2c"
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
		//	return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		//}
		return []byte("mysign"), nil
	})
	if err != nil {
		log.Println("error: ", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	log.Println(ok, claims)

}

func Test_Base64(t *testing.T) {
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njc3NTE3MDMsImZha2VFeHBpcmVzQXQiOjE1Njc0OTI1MTgsInRva2VuIjoidGVzdCJ9.vGDr16L8g75sXvNrxPWmHy5LkmOBagsFZjfziQTAqRg"
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("decode err: ", err)
		return
	}

	log.Printf("%s", data)
}

func Test_Bcrypt(t *testing.T) {
	passwd := []byte("123456")
	res, err := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("bcrypt err: ", err)
	}
	log.Printf("len=%d,password=%s", len(res), res)
	log.Println(bcrypt.CompareHashAndPassword(res, passwd))

	p := md5.Sum(passwd)
	pwd := p[:]

	res, err = bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("bcrypt err: ", err)
	}
	log.Printf("len=%d,password=%s", len(res), res)
	log.Println(bcrypt.CompareHashAndPassword(res, pwd))
}

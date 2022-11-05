package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var hash []byte

func main() {
	hash, _ = hashPassword("12")
	fmt.Println(string(hash))
	comparePassword("12")
	fmt.Println("")

	//hmac in action
	sign, _ := signMessage([]byte("arun"))
	checkSignature([]byte("arun"), sign)
}

func signMessage(msg []byte) ([]byte, error) {
	hash := hmac.New(sha512.New, []byte("12345"))
	_, err := hash.Write(msg)
	if err != nil {
		fmt.Println("Error while signing the message")
		return nil, err
	}
	signature := hash.Sum(nil)
	return signature, nil

}

func checkSignature(msg, sig []byte) {
	newSig, err := signMessage(msg)
	if err != nil {
		fmt.Println("Error signing message")
	}
	same := hmac.Equal(newSig, sig)
	fmt.Println(same)

}

func hashPassword(password string) ([]byte, error) {
	// fmt.Println(base64.StdEncoding.EncodeToString([]byte("username:pass")))
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func comparePassword(password string) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		fmt.Println("Wrong password")
	} else {
		fmt.Println("Success")
	}
}

package tokens

import (
	"fmt"
	"log"
	"testing"
)

type testCase struct {
	guidToken string
	mail      string
	ip        string
}

var testcase = &testCase{
	guidToken: "2234520a-abe8-4f60-90c8-0d43c5f6c0f6",
	mail:      "Ivanov@gmail.com",
	ip:        "198.27.164.1:8080",
}

func TestNewJWTToken(t *testing.T) {
	token, err := NewJWTToken(testcase.mail, testcase.ip)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println(token)
}

func TestParseJWT(t *testing.T) {
	token, err := NewJWTToken(testcase.mail, testcase.ip)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}

	mail, ip, err := ParseJWT(token)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	if ip == testcase.ip && mail == testcase.mail {
		log.Println("OK", ip, mail)
	} else {
		t.Fatalf("Ip or Mail is not the same\n")
	}
}

func TestMakeRefreshToken1(t *testing.T) {
	res, err := MakeRefreshToken(testcase.guidToken, testcase.ip)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}

	fmt.Println(res, 1)
}

func TestBase64(t *testing.T) {
	refresh := "$2a$13$WxXugNFsvEe7PQP02qbIFeP9Em5JZA.VgIBo2EusfttKy9zipGF4e"
	encd := EncodeBase(refresh)
	result, err := DecodeBase(encd)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}

	if refresh != result {
		t.Fatalf("original and decoded refresh tokens  are different")
		return
	}

	log.Println("OK")
	log.Println(refresh, "==", result)
}

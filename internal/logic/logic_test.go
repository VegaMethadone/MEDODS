package logic

import (
	"log"
	"testing"
)

type testCase struct {
	guidToken string
	ip        string
}

var testcase = &testCase{
	guidToken: "2234520e-abe8-4f60-90c8-0d43c5f6c0f6",
	ip:        "198.27.164.1",
}

func TestBusinessGetUserTokens(t *testing.T) {
	jwtToken, refreshToken, err := BusinessGetUserTokens(testcase.guidToken, testcase.ip)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println("JWT:", jwtToken)
	log.Println("REFRESH:", refreshToken)
}

func TestBusinessUpdateUserToken(t *testing.T) {
	refresh, _, err := database.Get(conf, testcase.guidToken)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}

	jwtToken, refreshToken, err := BusinessUpdateUserToken(testcase.guidToken, testcase.ip, refresh)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}

	log.Println(jwtToken)
	log.Println(refreshToken)
}

package postgres

import (
	"log"
	"medods/config"
	"medods/internal/tokens"
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

func TestAdd(t *testing.T) {
	refresh, err := tokens.MakeRefreshToken(testcase.guidToken, testcase.ip)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println(refresh)

	conf, _ := config.GetConfig()

	newPostgres := &Postgres{}
	err = newPostgres.Add(conf, testcase.guidToken, refresh, testcase.mail)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println("OK")
}

func TestCheck(t *testing.T) {
	conf, _ := config.GetConfig()
	newPostgres := &Postgres{}

	isOkay, err := newPostgres.Check(conf, testcase.guidToken)
	if !isOkay {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println("OK")
}

func TestUpdate(t *testing.T) {
	conf, _ := config.GetConfig()
	newPostgres := &Postgres{}

	refresh, err := tokens.MakeRefreshToken(testcase.guidToken, "127.0.0.1:8080")
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println(refresh)

	err = newPostgres.Update(conf, testcase.guidToken, refresh)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println("OK")
}

func TestGet(t *testing.T) {
	conf, _ := config.GetConfig()
	newPostgres := &Postgres{}

	refresh, mail, err := newPostgres.Get(conf, testcase.guidToken)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}

	log.Println(refresh)
	log.Println(mail)
	if mail == testcase.mail {
		log.Println("OK")
	}
}

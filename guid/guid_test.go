package guid

import (
	"fmt"
	"log"
	"testing"
)

var testCase string = "2234520a-abe8-4f60-90c8-0d43c5f6c0f6"

func TestValidateGUID(t *testing.T) {
	err := ValidateGUID(testCase)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println("OK")
}

func TestCreateGUID(t *testing.T) {
	newGuid, _ := CreateGUID()
	fmt.Println(newGuid)
	err := ValidateGUID(testCase)
	if err != nil {
		t.Fatalf("%v\n", err)
		return
	}
	log.Println("OK")
}

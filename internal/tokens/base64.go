package tokens

import (
	"encoding/base64"
)

func EncodeBase(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeBase(str string) (string, error) {
	original, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	return string(original), nil
}

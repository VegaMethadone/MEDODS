package tokens

import (
	"golang.org/x/crypto/bcrypt"
)

// просто беру для refresh token-a GUID,IP и делаю merge.
// Refresh токен тип произвольный,
func MakeRefreshToken(guidData, ip string) (string, error) {
	refreshStr := []byte{}
	l, r := 0, 0

	for l < len(guidData) || r < len(ip) {

		if l < len(guidData) {
			refreshStr = append(refreshStr, byte(guidData[l]))
			l++
		}

		if r < len(ip) {
			refreshStr = append(refreshStr, byte(ip[r]))
			r++
		}
	}

	// хранится в базе исключительно в виде bcrypt хеша
	refreshStr, err := bcrypt.GenerateFromPassword(refreshStr, 13)
	if err != nil {
		return "", err
	}

	return string(refreshStr), nil
}

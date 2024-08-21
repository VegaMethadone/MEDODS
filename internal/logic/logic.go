package logic

import (
	"errors"
	"fmt"
	"math/rand"
	"medods/config"
	"medods/internal/bd"
	"medods/internal/bd/postgres"
	"medods/internal/tokens"
)

var conf, _ = config.GetConfig()
var database = bd.NewDatabase(&postgres.Postgres{})
var chars string = "abcdefghijklmnopqrstuvwxyz"

// Первый маршрут выдает пару Access, Refresh токенов для пользователя с идентификатором (GUID) указанным в параметре запроса
func BusinessGetUserTokens(guidToken, ip string) (string, string, error) {

	/*	Я допустил тут 2 ситуации:
		1) GUID уже в бд -> тогда берет из бд  refresh token, делает  новый JWT token и выдает их юзеру
		2) GUID появляется в  первый раз -> создает рандомое мыло, создает из входных данных JWT и Refresh  токены, добавляет мыло, guid и refresh в бд и отправляет их юзеру
	*/

	exists, _ := database.Check(conf, guidToken)
	if !exists {
		mail := randEmail(20)
		// Генерация JWT токена
		jwtToke, err := tokens.NewJWTToken(mail, ip)
		if err != nil {
			return "", "", err
		}
		// Генерация Refresh токена
		refreshToken, err := tokens.MakeRefreshToken(guidToken, ip)
		if err != nil {
			return "", "", err
		}
		//	Добавление в бд данных
		err = database.Add(conf, guidToken, refreshToken, mail)
		if err != nil {
			return "", "", err
		}

		return jwtToke, refreshToken, nil
	}

	//	Получение существующего в бд Refresh токена
	refreshToken, mail, err := database.Get(conf, guidToken)
	if err != nil {
		return "", "", nil
	}
	// Генерация JWT токена
	jwtToken, err := tokens.NewJWTToken(mail, ip)
	if err != nil {
		return "", "", err
	}

	return jwtToken, refreshToken, nil
}

// Второй маршрут выполняет Refresh операцию на пару Access, Refresh токенов
func BusinessUpdateUserToken(guidToken, ip, refresh string) (string, string, error) {

	/*
		Тут, просто при получение куки и декодирование Refresh из base64 и  JWT
		Проверяется соотествие текущего Refresh из куки и Refresh из бд, если окей -> обновляем.  Не окей -> возвращаем  ошибку
		Проверяем старый и новый ip address пользователя проверяется в хендлере.
	*/

	exists, _ := database.Check(conf, guidToken)
	if !exists {
		return "", "", errors.New("no such user")
	}

	//	Получение нужных данных из бд  по GUID
	oldRefresh, mail, err := database.Get(conf, guidToken)
	if err != nil {
		return "", "", err
	}
	//	Проверка Refresh токенов на соответствие
	if oldRefresh != refresh {
		fmt.Println(oldRefresh, refresh)
		return "", "", errors.New("invalid refresh token")
	}

	//	Формирование нового JWT
	newJwtToken, err := tokens.NewJWTToken(mail, ip)
	if err != nil {
		return "", "", err
	}
	//	Формирование нового Refresh
	newRefreshToken, err := tokens.MakeRefreshToken(guidToken, ip)
	if err != nil {
		return "", "", err
	}

	// Попытка обновить данные в бд, если все ок -> вернет данные пользователю
	err = database.Update(conf, guidToken, newRefreshToken)
	if err != nil {
		return "", "", err
	}

	return newJwtToken, newRefreshToken, nil
}

func randEmail(n int) string {
	ends := []string{"@gmail.com", "@yandex.ru", "@mail.ru", "@list.ru"}
	fakeMailBody := make([]byte, n)

	for i := range n {
		fakeMailBody[i] = chars[rand.Intn(len(chars))]
	}

	return string(fakeMailBody) + ends[n%4]
}

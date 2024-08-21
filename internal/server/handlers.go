package server

import (
	"fmt"
	"log"
	"medods/guid"
	"medods/internal/logic"
	"medods/internal/tokens"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getUserTokens(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//	Извлекаю  GUID и ip
	path := mux.Vars(r)
	userGuid := path["guidToken"]
	addr := strings.Split(r.RemoteAddr, ":")

	//	Валидирую  GUID
	err := guid.ValidateGUID(userGuid)
	if err != nil {
		http.Error(w, "Invalid GUID", http.StatusBadRequest)
		return
	}

	//	Принимаю токены, которые были создани на уровне бизнес-логики
	jwtToke, refreshToken, err := logic.BusinessGetUserTokens(userGuid, addr[0])
	if err != nil {
		http.Error(w, "Faild to get tokens", http.StatusInternalServerError)
		return
	}

	//	Создаю пару куки
	jwtCookie := &http.Cookie{
		Name:     "medodstokenj",
		Value:    jwtToke,
		MaxAge:   60 * 60 * 24,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}
	refreshCookie := &http.Cookie{
		Name:     "medodstokenr",
		Value:    tokens.EncodeBase(refreshToken),
		MaxAge:   60 * 60 * 24,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}

	//	Устанавливаю  куки
	http.SetCookie(w, jwtCookie)
	http.SetCookie(w, refreshCookie)

	//	Все окей оби
	w.WriteHeader(http.StatusOK)
}

func updateUserTokens(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//	Извлекаю  GUID и ip
	path := mux.Vars(r)
	userGuid := path["guidToken"]
	addr := strings.Split(r.RemoteAddr, ":")

	//	Валидирую  GUID
	err := guid.ValidateGUID(userGuid)
	if err != nil {
		http.Error(w, "Invalid GUID", http.StatusBadRequest)
		return
	}

	//	Извлекаю все нужные куки
	jwtCookie, err := r.Cookie("medodstokenj")
	if err != nil {
		http.Error(w, "no valid medodstokenj cookie", http.StatusUnauthorized)
		return
	}
	refreshCookie, err := r.Cookie("medodstokenr")
	if err != nil {
		http.Error(w, "no valid medodstokenr cookie", http.StatusUnauthorized)
		return
	}

	//	Обрабатываю JWT токен
	oldEmail, addrFromJwt, err := tokens.ParseJWT(jwtCookie.Value)
	if err != nil {
		http.Error(w, "invalid jwt token", http.StatusUnauthorized)
		return
	}

	//	Тут работаю с текущем и старым адресом, сравниваю их: если ок -> просто дальше иду по коду
	oldAdd := strings.Split(addrFromJwt, ":")
	if addr[0] != oldAdd[0] {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(fmt.Sprintf("User has not the same ip: send a warning to %s", oldEmail)))
		return
	}

	//	Обновляю токены
	refreshToken, err := tokens.DecodeBase(refreshCookie.Value)
	if err != nil {
		log.Println(userGuid, err)
		http.Error(w, "cookie is damadged", http.StatusUnauthorized)
		return
	}
	jwtToken, refreshToken, err := logic.BusinessUpdateUserToken(userGuid, addr[0], refreshToken)
	if err != nil {
		log.Println(userGuid, err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// Переписываю в полученых куки поля Value, другие парметры остаются теми же
	jwtCookie.Value = jwtToken
	refreshCookie.Value = tokens.EncodeBase(refreshToken)

	//	Устанавливаю  куки
	http.SetCookie(w, jwtCookie)
	http.SetCookie(w, refreshCookie)

	//	Все окей оби
	w.WriteHeader(http.StatusOK)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Просто для проверки сервера
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("modeds home page!"))
}

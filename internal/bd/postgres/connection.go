package postgres

import (
	"errors"
	"fmt"
	"medods/config"
)

// возвращает строку и  ошибку для подключения к бд, которая берет данные из конфига
func connectionString(conf *config.Config) (string, error) {
	if conf == nil {
		return "", errors.New("could not find db config")
	}
	str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.DatabaseName, conf.Postgres.Sslmode, conf.Postgres.Host)

	return str, nil
}

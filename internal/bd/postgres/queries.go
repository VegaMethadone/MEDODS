package postgres

import (
	"database/sql"
	"medods/config"

	_ "github.com/lib/pq"
)

type Postgres struct{}

/*
Структура  БД

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    guid VARCHAR(256),
    refresh VARCHAR(256),
    email VARCHAR(256)
);
*/

func (p *Postgres) Check(conf *config.Config, guidData string) (bool, error) {
	//	Получаю строку для подключения к бд
	str, err := connectionString(conf)
	if err != nil {
		return false, err
	}
	//	Открываю соединение с бд
	db, err := sql.Open("postgres", str)
	if err != nil {
		return false, err
	}
	//	Обязательно его закрываю, чтобы завершить системный вызов
	defer db.Close()

	//	Выполняю запрос
	var count int
	err = db.QueryRow(
		`SELECT COUNT(*) FROM  users WHERE guid = $1`,
		guidData,
	).Scan(&count)

	if err != nil {
		return false, nil
	}

	return count > 0, nil
}

func (p *Postgres) Add(conf *config.Config, guidData, refresh, mail string) error {
	//	Получаю строку для подключения к бд
	str, err := connectionString(conf)
	if err != nil {
		return err
	}
	//	Открываю соединение с бд
	db, err := sql.Open("postgres", str)
	if err != nil {
		return err
	}
	//	Обязательно его закрываю, чтобы завершить системный вызов
	defer db.Close()

	//	Выполняю запрос
	_, err = db.Exec(
		`INSERT INTO users (guid, refresh, email) VALUES ($1, $2, $3)`,
		guidData, refresh, mail,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Update(conf *config.Config, guidData, refresh string) error {
	//	Получаю строку для подключения к бд
	str, err := connectionString(conf)
	if err != nil {
		return err
	}
	//	Открываю соединение с бд
	db, err := sql.Open("postgres", str)
	if err != nil {
		return err
	}
	//	Обязательно его закрываю, чтобы завершить системный вызов
	defer db.Close()

	//	Выполняю запрос
	_, err = db.Exec(
		`UPDATE users SET refresh = $1 WHERE guid = $2`,
		refresh, guidData,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Get(conf *config.Config, guidData string) (string, string, error) {
	//	Получаю строку для подключения к бд
	str, err := connectionString(conf)
	if err != nil {
		return "", "", err
	}
	//	Открываю соединение с бд
	db, err := sql.Open("postgres", str)
	if err != nil {
		return "", "", err
	}
	//	Обязательно его закрываю, чтобы завершить системный вызов
	defer db.Close()

	//	Выполняю запрос
	var refresh, mail string
	err = db.QueryRow(
		`SELECT refresh, email FROM users WHERE guid = $1`,
		guidData,
	).Scan(&refresh, &mail)

	if err != nil {
		return "", "", err
	}

	return refresh, mail, nil
}

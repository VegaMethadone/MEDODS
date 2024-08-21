package bd

import (
	"medods/config"
	"medods/internal/bd/postgres"
)

/*

Тут я реализовал паттерн проектирования  - состояние. Если для  выдачи токенов и т.д. понадобится другаябд,
можно просто создать отдельный пакет, реализовать все функции из интерфейса и переключать состояние
на нужную бд, если это необходимо

*/

type Database struct {
	postgres State
	//Mongo  State
	//Redis  State

	current State
}

type State interface {
	Add(*config.Config, string, string, string) error
	Get(*config.Config, string) (string, string, error)
	Update(*config.Config, string, string) error
	Check(*config.Config, string) (bool, error)
}

func NewDatabase(ps *postgres.Postgres) *Database {
	return &Database{
		postgres: ps,
		current:  ps,
	}
}

func (db *Database) SetState(state State) {
	db.current = state
}

func (db *Database) Add(conf *config.Config, guidData, refresh, mail string) error {
	return db.current.Add(conf, guidData, refresh, mail)
}
func (db *Database) Get(conf *config.Config, guidData string) (string, string, error) {
	return db.current.Get(conf, guidData)
}
func (db *Database) Update(conf *config.Config, guidData, refresh string) error {
	return db.current.Update(conf, guidData, refresh)
}
func (db *Database) Check(conf *config.Config, guidData string) (bool, error) {
	return db.current.Check(conf, guidData)
}

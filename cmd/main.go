package main

import (
	"fmt"
	"log"
	"medods/config"
	"medods/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//	Извлекаю  конфиг для дальнейшей работы с ним
	conf, err := config.GetConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//	Создаю новый сервер  на основе моего конфига
	srv := server.NewServer(conf)
	fmt.Println("Server is working at: http://127.0.0.1:8080/home")

	//	Создаю  каналы для работы  с ошибкой от сервера и принудительным выключением от хоста
	done := make(chan os.Signal, 1)
	waitingForErr := make(chan bool, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	//	Запускаю сервер в отдельной горутине
	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
			waitingForErr <- true
		}
	}()

	//	Создаю бесконечный цикл до первой ошибки или принудительного выключения
	select {
	case s := <-done:
		log.Println("Server is stoped by host", s)
		srv.Close()
		waitingForErr <- true
	case err := <-waitingForErr:
		log.Println("Error occured:", err)
	}

	log.Println("Server is shut down")
}

package main

import (
	"Tasks_Users_Vk_test/internal/repository"
	"Tasks_Users_Vk_test/internal/service"
	"Tasks_Users_Vk_test/internal/transport"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// @title User_Quests API

func main() {

	// загружаем файл конфигурации
	viper.SetConfigFile("config/config.yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Получаем значения из конфигурации
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	// используем данные из файла конфигурации для подключения к бд
	repos, err := repository.NewRepositories(dbname, username, password, host, port)

	servs, err := service.NewServices(repos)

	handlerUser := transport.NewUserHandler(repos)
	handlerQuest := transport.NewQuestHandler(repos)
	handlerCompletedTask := transport.NewCompletedQuestsHandler(repos, servs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server started...")

	router := mux.NewRouter()

	router.HandleFunc("/users/{id:[0-9]+}", handlerUser.GetUser).Methods("GET")
	router.HandleFunc("/users", handlerUser.CreateUser).Methods("POST")
	router.HandleFunc("/quest", handlerQuest.CreateQuest).Methods("POST")
	router.HandleFunc("/quest/{id:[0-9]+}", handlerQuest.GetQuest).Methods("GET")
	router.HandleFunc("/complete", handlerCompletedTask.CompleteTask).Methods("POST")
	router.HandleFunc("/complete/{id:[0-9]+}", handlerCompletedTask.GetCompletedQuestsAndBalance).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

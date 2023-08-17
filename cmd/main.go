package main

import (
	"github.com/hurucik/TODO_app"
	"github.com/hurucik/TODO_app/pckg/handler"
	"github.com/hurucik/TODO_app/pckg/repository"
	"github.com/hurucik/TODO_app/pckg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initialazing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repose := repository.NewRepository(db)
	services := service.NewService(repose)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run((viper.GetString("8000")), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

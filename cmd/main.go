package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/mxmntv/todo_app/pkg/repository"
	"github.com/mxmntv/todo_app/pkg/service"

	todo "github.com/mxmntv/todo_app"
	"github.com/mxmntv/todo_app/pkg/handler"
)

func init() {
	if err := initConfig(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	server := new(todo.Server)
	if err := server.Start(viper.GetString("port"), handler.InitRoutes()); err != nil {
		fmt.Println("err widt starting server")
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

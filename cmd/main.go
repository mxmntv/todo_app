package main

import (
	"fmt"

	todo "github.com/mxmntv/todo_app"
	"github.com/mxmntv/todo_app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	server := new(todo.Server)
	if err := server.Start("8000", handler.InitRoutes()); err != nil {
		fmt.Println("err widt starting server")
	}
}

package main

import (
	// "github.com/go-sql-driver/mysql"
	// "github.com/ntphiep/go-ecom/cmd/api"
	// "github.com/ntphiep/go-ecom/config"
	// "github.com/ntphiep/go-ecom/db"

	"github.com/gorilla/mux"
	"github.com/ntphiep/go-ecom/pkg/handler"

	"log"
	"net/http"
)

func main() {
	// db, err := db.NewMySQLStorage(mysql.Config{
	// 	User:                 config.Envs.DBUser,
	// 	Passwd:               config.Envs.DBPassword,
	// 	Addr:                 config.Envs.DBAddress,
	// 	DBName:               config.Envs.DBName,
	// 	Net:                  "tcp",
	// 	AllowNativePasswords: true,
	// 	ParseTime:            true,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server := api.NewApiServer(":8080", nil)
	// if err := server.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	r := mux.NewRouter()

	r.HandleFunc("/api/todo", handler.GetAllTodo).Methods(http.MethodGet)	
	r.HandleFunc("/api/todo/{id}", handler.GetTodoById).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id}", handler.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id}", handler.DeleteTodo).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}

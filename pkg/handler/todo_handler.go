package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/ntphiep/go-ecom/pkg/data"
	"github.com/ntphiep/go-ecom/pkg/dto"
)

func GetAllTodo(writer http.ResponseWriter, req *http.Request) {
	responseWithJSON(writer, http.StatusOK, data.Todos)
}

func GetTodoById(writer http.ResponseWriter, request *http.Request) {

}

func CreateTodo(writer http.ResponseWriter, request *http.Request) {

}

func UpdateTodo(writer http.ResponseWriter, request *http.Request) {

}

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	
}


func responseWithJSON(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func generateId(todos []dto.Todo) int {
	var maxId int
	for _, todo := range todos {
		if todo.ID > maxId {
			maxId = todo.ID
		}
	}

	return maxId + 1
}
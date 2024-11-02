package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ntphiep/go-ecom/pkg/data"
	"github.com/ntphiep/go-ecom/pkg/dto"
)

func GetAllTodo(writer http.ResponseWriter, req *http.Request) {
	responseWithJSON(writer, http.StatusOK, data.Todos)
}

func GetTodoById(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, "Invalid ID")
		return
	}

	for _, todo := range data.Todos {
		if todo.ID == id {
			responseWithJSON(writer, http.StatusOK, todo)
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, "Todo not found")
}

func CreateTodo(writer http.ResponseWriter, req *http.Request) {
	var newTodo dto.Todo
	if err := json.NewDecoder(req.Body).Decode(&newTodo); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	newTodo.ID = generateId(data.Todos)
	data.Todos = append(data.Todos, newTodo)

	responseWithJSON(writer, http.StatusCreated, newTodo)
}

func UpdateTodo(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	var updateTodo dto.Todo
	if err := json.NewDecoder(req.Body).Decode(&updateTodo); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateTodo.ID = id

	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos[i] = updateTodo
			responseWithJSON(writer, http.StatusOK, updateTodo)
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func DeleteTodo(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
			responseWithJSON(writer, http.StatusOK, map[string]string{"message": "Todo was deleted"})
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
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

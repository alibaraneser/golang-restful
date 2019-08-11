package controllers

import (
	"encoding/json"
	"golang-todo/models"
	u "golang-todo/utils"
	"net/http"
)




type GetTodoViewModel struct{
		ID uint `json:"id"`
}

var CreateTodo = func(w http.ResponseWriter, r *http.Request) {

	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := todo.Create()
	u.Respond(w, resp)
}

var GetTodo = func(w http.ResponseWriter, r *http.Request) {

	info := &GetTodoViewModel{}

	err := json.NewDecoder(r.Body).Decode(info)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := models.GetTodo(info.ID)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)

}
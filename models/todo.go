package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "golang-todo/utils"
)

type Todo struct {
	gorm.Model
	Title string `json:"title",gorm:"not null"`
	UserId uint `json:"user_id"`
}

func (todo *Todo) Validate() (map[string] interface{}, bool) {
	fmt.Print(todo)
	if todo.Title == "" {
		return u.Message(false, "Title is required!"), false
	}

	if todo.UserId <= 0 {
		return u.Message(false, "User is not defined"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (todo *Todo) Create() (map[string] interface{}) {

	if resp, ok := todo.Validate(); !ok {
		return resp
	}

	GetDB().Create(todo)

	resp := u.Message(true, "success")
	resp["todo"] = todo
	return resp
}

func GetTodo(id uint) (*Todo) {

	todo := &Todo{}
	err := GetDB().Table("todos").Where("id = ?", id).First(todo).Error
	if err != nil {
		return nil
	}
	return todo
}

func GetTodos(user uint) ([]*Todo) {

	todos := make([]*Todo, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&todos).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return todos
}


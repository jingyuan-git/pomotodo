package service

import (
	"fmt"
	"log"

	"server/models"
)

type Todo struct {
	ID          string `json:"id,omitempty" gorm:"primarykey"`
	Title       string `json:"title,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	IsCompleted bool   `json:"isCompleted,omitempty"`

	FilterBegin string `json:"filterBegin,omitempty"`
	FilterEnd   string `json:"filterEnd,omitempty"`
}

type TodoDto struct {
	ID    string `json:"id" gorm:"primarykey"`
	Title string `json:"title,omitempty"`
}

func (a *Todo) CreateTodo() error {
	createTime, _ := formatTimeStamp(a.CreatedAt)

	todo := models.Todo{
		ID:          a.ID,
		Title:       a.Title,
		CreatedAt:   createTime,
		IsCompleted: a.IsCompleted,
	}
	fmt.Printf("service %+v/n", todo)
	if err := models.CreateTodo(todo); err != nil {
		return err
	}

	return nil
}

func (a *Todo) GetAll() ([]*models.Todo, error) {
	aa := make(map[string]interface{})
	todos, err := models.GetTodos(aa)
	if err != nil {
		log.Default().Printf("fail to list all pomos, error: %+v \n", err)
		return nil, err
	}

	return todos, nil
}

func (a *Todo) DeleteTodo() error {
	createTime, _ := formatTimeStamp(a.CreatedAt)

	todo := models.Todo{
		ID:          a.ID,
		Title:       a.Title,
		CreatedAt:   createTime,
		IsCompleted: a.IsCompleted,
	}
	fmt.Printf("service %+v/n", todo)
	if err := models.DeleteTodo(todo); err != nil {
		return err
	}

	return nil
}

func (a *Todo) UpdateTodo() error {
	createTime, _ := formatTimeStamp(a.CreatedAt)

	todo := models.Todo{
		ID:          a.ID,
		Title:       a.Title,
		CreatedAt:   createTime,
		IsCompleted: a.IsCompleted,
	}
	fmt.Printf("service %+v \n", todo)
	if err := models.UpdateTodo(todo); err != nil {
		return err
	}

	return nil
}

func (a *Todo) CountAndTotalAmount() (int, error) {

	return 0, nil
}

func (a *Todo) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	filterBegin, _ := formatTimeStamp(a.FilterBegin)
	filterEnd, _ := formatTimeStamp(a.FilterEnd)

	if a.FilterBegin != "" {
		maps["filterBegin"] = filterBegin
	}
	if a.FilterEnd != "" {
		maps["filterEnd"] = filterEnd
	}
	return maps
}

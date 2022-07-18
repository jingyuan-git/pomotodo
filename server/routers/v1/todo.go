package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/pkg/app"
	"server/pkg/e"

	// "server/pkg/setting"
	// "server/pkg/util"
	// "server/models"
	"server/service"
)

func GetTodos(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form service.Todo
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	todoService := service.Todo{}

	total, err := todoService.CountAndTotalAmount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	todos, err := todoService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = todos
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func CreateTodo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	form := service.Todo{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoService := service.Todo{
		ID:          form.ID,
		Title:       form.Title,
		CreatedAt:   form.CreatedAt,
		IsCompleted: form.IsCompleted,
	}

	err := todoService.CreateTodo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteTodo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	form := service.Todo{}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoService := service.Todo{
		ID: form.ID,
	}

	err := todoService.DeleteTodo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func UpdateTodo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	form := service.Todo{}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoService := service.Todo{
		ID:          form.ID,
		Title:       form.Title,
		CreatedAt:   form.CreatedAt,
		IsCompleted: form.IsCompleted,
	}

	err := todoService.UpdateTodo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

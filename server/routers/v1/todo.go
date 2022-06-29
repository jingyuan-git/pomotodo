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

// func CountTodos(c *gin.Context) {
// 	var (
// 		appG = app.Gin{C: c}
// 		form service.Todo
// 	)

// 	// httpCode, errCode := app.BindAndValid(c, &form)
// 	// if errCode != e.SUCCESS {
// 	// 	appG.Response(httpCode, errCode, nil)
// 	// 	return
// 	// }

// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	todoService := service.Todo{
// 		FilterBegin: form.FilterBegin,
// 		FilterEnd:   form.FilterEnd,
// 	}

// 	fmt.Println("todoService", todoService, "form", form)

// 	total, amount, err := todoService.CountAndTotalAmount()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
// 		return
// 	}

// 	orders, err := todoService.CountTodos()
// 	if err != nil {
// 		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
// 		return
// 	}

// 	data := make(map[string]interface{})
// 	data["lists"] = orders
// 	data["total"] = total
// 	data["totalAmount"] = amount

// 	appG.Response(http.StatusOK, e.SUCCESS, data)
// }

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

	orders, err := todoService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = orders
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

	data := make(map[string]interface{})
	// data["lists"] = orders
	// data["total"] = total
	// data["totalAmount"] = amount

	appG.Response(http.StatusOK, e.SUCCESS, data)
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

	data := make(map[string]interface{})
	// data["lists"] = orders
	// data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

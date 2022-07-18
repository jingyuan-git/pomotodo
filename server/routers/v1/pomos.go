package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"server/pkg/app"
	"server/pkg/e"

	// "server/pkg/setting"
	// "server/pkg/util"
	// "server/models"
	"server/service"
)

func CountPomos(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form service.Pomo
	)

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pomoService := service.Pomo{
		ID:        form.ID,
		Title:     form.Title,
		StartTime: form.StartTime,
		EndTime:   form.EndTime,

		FilterBegin: form.FilterBegin,
		FilterEnd:   form.FilterEnd,
	}

	fmt.Println("pomoService", pomoService, "form", form)

	pomos, err := pomoService.CountPomos()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = pomos
	data["total"] = len(pomos)

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetPomos(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form service.Pomo
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	pomoService := service.Pomo{
		ID:        form.ID,
		Title:     form.Title,
		StartTime: form.StartTime,
		EndTime:   form.EndTime,

		FilterBegin: form.FilterBegin,
		FilterEnd:   form.FilterEnd,
	}

	pomos, err := pomoService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = pomos
	data["total"] = len(pomos)

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func CreatePomo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	form := service.Pomo{}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pomoService := service.Pomo{
		ID:        form.ID,
		Title:     form.Title,
		StartTime: form.StartTime,
		EndTime:   form.EndTime,

		FilterBegin: form.FilterBegin,
		FilterEnd:   form.FilterEnd,
	}

	err := pomoService.CreatePomo()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

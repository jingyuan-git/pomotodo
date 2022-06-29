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

	// httpCode, errCode := app.BindAndValid(c, &form)
	// if errCode != e.SUCCESS {
	// 	appG.Response(httpCode, errCode, nil)
	// 	return
	// }

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pomoService := service.Pomo{
		FilterBegin: form.FilterBegin,
		FilterEnd:   form.FilterEnd,
	}

	fmt.Println("pomoService", pomoService, "form", form)

	total, amount, err := pomoService.CountAndTotalAmount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	orders, err := pomoService.CountPomos()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = orders
	data["total"] = total
	data["totalAmount"] = amount

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

	pomoService := service.Pomo{}

	total, amount, err := pomoService.CountAndTotalAmount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	orders, err := pomoService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = orders
	data["total"] = total
	data["totalAmount"] = amount

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func CreatePomo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	form := service.PomoDto{}
	fmt.Println("00000000000")
	// if err := c.BindJSON(&form); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("form %+v \n", form)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// httpCode, errCode := app.BindAndValid(c, &form)
	fmt.Println("111111111111111")
	fmt.Printf("routers  %+v/n", form)
	// fmt.Println(httpCode, errCode)
	// if errCode != e.SUCCESS {
	// 	appG.Response(httpCode, errCode, nil)
	// 	return
	// }

	pomoService := service.Pomo{}

	err := pomoService.CreatePomo(form)
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

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"entitys"
	"models"
	"utils"

	"github.com/labstack/echo"
)

type handler struct {
	db map[string]*entitys.V_employee_profile
}

func GetAll(c echo.Context) error {
	data, err := models.All()

	var content []map[string]interface{}
	json_byte, _ := json.Marshal(data)
	json.Unmarshal(json_byte, &content)

	body := map[string]interface{}{
		"data": content,
	}

	resp, _ := utils.Response_Handle(body, err)
	return c.JSON(http.StatusOK, resp)
}

func Get(c echo.Context) error {
	number, _ := strconv.ParseInt(c.Param("count"), 0, 64)

	data, err := models.Count(number)

	var content []map[string]interface{}
	json_byte, _ := json.Marshal(data)
	json.Unmarshal(json_byte, &content)

	body := map[string]interface{}{
		"data": content,
	}

	resp, _ := utils.Response_Handle(body, err)
	return c.JSON(http.StatusOK, resp)
}

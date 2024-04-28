package handlers

import (
	"net/http"
	"strconv"

	"github.com/dcrespo1/mimir/cmd/api/services"
	"github.com/labstack/echo"
)

func PostIndexHandler(c echo.Context) error {
	data, err := services.GetAll()
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to Process Data")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to process Data")
	}
	data, err := services.GetById(idx)
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to Process Data")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func list(c echo.Context) error {
	data := map[string]interface{}{}
	sm.Range(func(k, v interface{}) bool {
		data[k.(string)] = v
		return true
	})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func create(c echo.Context) error {
	var body map[string]interface{}
	json.NewDecoder(c.Request().Body).Decode(&body)

	for key, val := range body {
		sm.Store(key, val)
	}

	data := map[string]interface{}{}
	sm.Range(func(k, v interface{}) bool {
		data[k.(string)] = v
		return true
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func delete(c echo.Context) error {
	sm.Delete(c.Param("key"))
	data := map[string]interface{}{}
	sm.Range(func(k, v interface{}) bool {
		data[k.(string)] = v
		return true
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zombiemachines/echo-templ-htmx-tailwindcss/models"
	"github.com/zombiemachines/echo-templ-htmx-tailwindcss/views"
)

func HomeHandler(c echo.Context) error {
	context := c.Request().Context()
	writer := c.Response().Writer
	return views.Hello("Lyoko").Render(context, writer)
}

func HelloPostHandler(c echo.Context) error {
	context := c.Request().Context()
	writer := c.Response().Writer
	var v = models.Visitor{LikesDueling: true}
	v.Name = c.FormValue("name")
	if v.Name != "" {
		return views.Card(v.Name).Render(context, writer)

	}
	return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid name")
}

func FormHandler(c echo.Context) error {
	context := c.Request().Context()
	writer := c.Response().Writer
	return views.FormHello().Render(context, writer)
}

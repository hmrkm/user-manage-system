package main

import (
	"github.com/hmrkm/user-manage-system/adapter"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, ua adapter.Users, aa adapter.Auth) {
	e.POST("/v1/auth", func(c echo.Context) error {
		req := &adapter.RequestAuth{}
		if err := c.Bind(req); err != nil {
			return c.JSON(400, nil)
		}

		res, err := aa.Auth(c.Request().Context(), req)
		if err != nil {
			return ErrorHandler(c, err)
		}

		return c.JSON(200, res)
	})
	e.POST("/v1/users/list", func(c echo.Context) error {
		req := &adapter.RequestUsersList{}
		if err := c.Bind(req); err != nil {
			return c.JSON(400, nil)
		}

		res, err := ua.List(c.Request().Context(), req)
		if err != nil {
			return ErrorHandler(c, err)
		}

		return c.JSON(200, res)
	})
}

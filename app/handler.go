package main

import (
	"github.com/hmrkm/user-manage-system/domain"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(c echo.Context, err error) error {
	switch err {
	case domain.ErrNotFound:
		return c.JSON(404, nil)
	case domain.ErrNotAuthorized:
		return c.JSON(401, nil)
	case domain.ErrForbidden:
		return c.JSON(403, nil)
	default:
		return c.JSON(500, nil)
	}
}

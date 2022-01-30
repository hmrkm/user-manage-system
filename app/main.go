package main

import (
	"github.com/hmrkm/user-manage-system/adapter"
	"github.com/hmrkm/user-manage-system/io"
	"github.com/hmrkm/user-manage-system/usecase"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

func main() {
	config := Config{}
	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}

	http := io.NewHTTP(5, 5)
	au := usecase.NewAuth(http, config.AuthEndpoint)
	uu := usecase.NewUsers(http, config.VerifyEndpoint, config.UserManageBaseurl)
	ua := adapter.NewUsers(uu)
	aa := adapter.NewAuth(au)

	e := echo.New()
	Router(e, ua, aa)

	e.Logger.Fatal(e.Start(":80"))
}

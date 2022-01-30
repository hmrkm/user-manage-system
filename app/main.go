package main

import (
	"github.com/hmrkm/user-manage-system/adapter"
	"github.com/hmrkm/user-manage-system/io"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

func main() {
	config := Config{}
	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}

	http := io.NewHTTP(5, 5)
	ua := adapter.NewUserAdapter(http, config.AuthEndpoint)
	aa := adapter.NewAuthAtapdater(http)

	e := echo.New()
	Router(e, ua, aa)

	e.Logger.Fatal(e.Start(":80"))
}

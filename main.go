package main

import (
	"echogorm1/business"
	"echogorm1/config"
	"echogorm1/model"
	"echogorm1/presentation"
	"github.com/labstack/echo/v4"
)

func main() {

	db := config.ConnectDB()
	db.AutoMigrate(&model.User{})

	userRepo := &model.UserRepo{db}

	guest := business.Guest{
		GetUserByEmail: userRepo.GetUserByEmail,
		InsertUser:     userRepo.InsertUser,
	}

	e := echo.New()

	e.POST(business.Guest_LoginRoute, func(c echo.Context) error {
		in := business.Guest_LoginIn{}
		if err := presentation.ParseInput(c, &in, &in.CommonRequest); err != nil {
			return err
		}
		out := guest.Login(&in)
		return presentation.RenderOutput(c, &out, &out.CommonResponse)
	})

	e.POST(business.Guest_RegisterRoute, func(c echo.Context) error {
		in := business.Guest_RegisterIn{}
		if err := presentation.ParseInput(c, &in, &in.CommonRequest); err != nil {
			return err
		}
		out := guest.Register(&in)
		return presentation.RenderOutput(c, &out, &out.CommonResponse)
	})

	e.Logger.Fatal(e.Start(config.ListenAddr))
}

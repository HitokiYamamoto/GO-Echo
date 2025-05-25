package router

import (
	"go-echo/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	// デバッグ用: ルーティング確認
	for _, r := range e.Routes() {
		println("Route:", r.Method, r.Path)
	}

	return e
}

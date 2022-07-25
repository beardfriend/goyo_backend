package server

import (
	"goyo/modules/academy"
	"goyo/modules/common"
	"goyo/modules/health"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	controllers := []common.Controller{
		health.HealthController{},
		academy.AcademyController{},
	}

	var routes []common.Route

	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}

	api := e.Group("")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
				break
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
				break
			}
		}
	}
}

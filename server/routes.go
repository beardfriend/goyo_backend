package server

import (
	"log"

	"goyo/modules/academy"
	"goyo/modules/common"
	"goyo/modules/health"
	"goyo/modules/yoga"

	"github.com/gin-gonic/gin"
)

func GinRoutes(engine *gin.Engine) {
	if err := engine.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	controllers := []common.Controller{
		health.HealthController{},
		academy.AcademyController{},
		yoga.YogaController{},
	}

	var routes []common.Route

	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}

	api := engine.Group("/api")

	for _, route := range routes {
		switch route.Method {

		case "POST":
			{

				api.POST(route.Path, route.Handler...)
				break
			}
		case "GET":
			{
				api.GET(route.Path, route.Handler...)
				break
			}
		case "DELETE":
			{
				api.DELETE(route.Path, route.Handler...)
				break
			}
		case "PUT":
			{
				api.PUT(route.Path, route.Handler...)
				break
			}
		case "PATCH":
			{
				api.PATCH(route.Path, route.Handler...)
				break
			}
		}
	}
}

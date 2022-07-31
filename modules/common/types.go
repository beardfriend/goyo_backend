package common

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler []gin.HandlerFunc
}
type Controller interface {
	Routes() []Route
}

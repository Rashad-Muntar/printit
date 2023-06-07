package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rashad-muntar/printit/config"
	"github.com/rashad-muntar/printit/routes"
)

func init() {
	config.LoadInitializers()
	config.ConnectDB()
}

func main() {
	r := gin.New()
	routes.AuthRoute(r)
	r.Run(":8080")
}

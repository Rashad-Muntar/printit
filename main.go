package main

import (
	"github.com/rashad-muntar/printit/routes"
	"github.com/gin-gonic/gin"
	)

func init() {

}

func main() {
	r := gin.New()
	routes.AuthRoute(r)
	r.Run(":8080")
}

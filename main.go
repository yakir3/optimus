package main

import (
	"go-pic-tools/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// load html template
	r.LoadHTMLGlob("templates/*")
	// load static directory
	r.Static("/static", "./static")

	// route group
	route.InitRoute(r)

	r.Run()
}

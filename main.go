package main

import (
	"github.com/SakethAjith/RESTfulBlog/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetRoutes(r)
	r.Run(":8080")
}

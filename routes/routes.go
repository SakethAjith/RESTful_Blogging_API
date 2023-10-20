package routes

import (
	"github.com/SakethAjith/RESTfulBlog/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	blogs := r.Group("/blogs")
	{
		blogs.GET("/", handlers.GetBlogs)
		blogs.GET("/:id", handlers.GetBlog)
		blogs.POST("/", handlers.CreateBlog)
		blogs.PUT("/:id", handlers.UpdateBlog)
		blogs.DELETE("/:id", handlers.DeleteBlog)
	}
}

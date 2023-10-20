package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SakethAjith/RESTfulBlog/database"
	"github.com/SakethAjith/RESTfulBlog/models"
	"github.com/gin-gonic/gin"
)

func GetBlogs(c *gin.Context) {
	// Implement logic to fetch all blog Blogs from the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	Blogs := []models.Blogs{}
	if err := db.Select(&Blogs, "SELECT * FROM blogs"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve any blog"})
		return
	}

	c.JSON(http.StatusOK, Blogs)
}

func GetBlog(c *gin.Context) {
	// Implement logic to fetch a specific blog Blog by ID from the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	Blogs := models.Blogs{}
	if err := db.Get(&Blogs, "SELECT * FROM blogs WHERE id = $1", id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, Blogs)
}

func CreateBlog(c *gin.Context) {
	// Implement logic to create a new blog Blog in the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	var Blogs models.Blogs

	if err := c.BindJSON(&Blogs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	fmt.Println(Blogs)
	result, err := db.NamedExec("INSERT INTO blogs VALUES (DEFAULT,:title, :content)", Blogs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new blog"})
		return
	}

	Blogs.Id, _ = result.LastInsertId()
	c.IndentedJSON(http.StatusCreated, Blogs)
}

func UpdateBlog(c *gin.Context) {
	// Implement logic to update an existing blog Blog in the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	var Blogs models.Blogs

	if err := c.BindJSON(&Blogs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	Blogs.Id = id

	_, err = db.NamedExec("UPDATE blogs SET title=:title, content=:content WHERE id=:id", Blogs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the blog"})
		return
	}
	c.IndentedJSON(http.StatusOK, Blogs)

}

func DeleteBlog(c *gin.Context) {
	// Implement logic to delete a blog Blog by ID from the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	_, err = db.Exec("DELETE FROM blogs WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the blog"})
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}

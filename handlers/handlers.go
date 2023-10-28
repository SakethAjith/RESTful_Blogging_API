package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SakethAjith/RESTfulBlog/database"
	"github.com/SakethAjith/RESTfulBlog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func validId(id int64, db *gorm.DB, c *gin.Context) bool {
	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return false
	}
	var blogs models.Blogs
	check := db.Where("id=?", id).First(&blogs)

	errors.Is(check.Error, gorm.ErrRecordNotFound)
	if check.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blog with that Id does not exist"})
		return false
	}

	return true
}

func GetBlogs(c *gin.Context) {
	// Implement logic to fetch all blog Blogs from the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	Blogs := []models.Blogs{}
	db.Find(&Blogs)

	c.JSON(http.StatusOK, Blogs)
}

func GetBlog(c *gin.Context) {
	// Implement logic to fetch a specific blog Blog by ID from the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	if !validId(id, db, c) {
		return
	}

	Blogs := models.Blogs{}
	db.Where("id=?", id).Find(&Blogs)
	c.JSON(http.StatusOK, Blogs)
}

func CreateBlog(c *gin.Context) {
	// Implement logic to create a new blog Blog in the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var Blogs models.Blogs

	if err := c.BindJSON(&Blogs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db.Create(&Blogs)

	db.Last(&Blogs)
	c.IndentedJSON(http.StatusCreated, Blogs)
}

func UpdateBlog(c *gin.Context) {
	// Implement logic to update an existing blog Blog in the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	if !validId(id, db, c) {
		return
	}

	var Blogs models.Blogs

	if err := c.BindJSON(&Blogs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var initBlog models.Blogs
	db.Where("id = ?", id).Find(&initBlog)
	initBlog.Title = Blogs.Title
	initBlog.Content = Blogs.Content
	db.Save(&initBlog)
	Blogs = initBlog

	c.IndentedJSON(http.StatusOK, Blogs)

}

func DeleteBlog(c *gin.Context) {
	// Implement logic to delete a blog Blog by ID from the database
	db, err := database.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog Id"})
		return
	}

	if !validId(id, db, c) {
		return
	}

	db.Delete(&models.Blogs{}, id)

	c.JSON(http.StatusNoContent, gin.H{"info": "Deleted Blog successfully!!"})
}

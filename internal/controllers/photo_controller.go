package controllers

import (
	"net/http"
	"task-5-pbi-btpns-deianearra/internal/config/database"
	"task-5-pbi-btpns-deianearra/internal/models"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var newPhoto models.Photo
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.DB
	if err := db.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Photo created successfully"})
}

func GetAllPhotos(c *gin.Context) {
	var photos []models.Photo

	db := database.DB
	if err := db.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching photos"})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	var updatedPhoto models.Photo

	db := database.DB
	if err := db.Where("id = ?", photoID).First(&updatedPhoto).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&updatedPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	var photo models.Photo

	db := database.DB
	if err := db.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := db.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

// func CreatePhoto(c *gin.Context) {
// 	var newPhoto models.Photo
// 	if err := c.ShouldBindJSON(&newPhoto); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Create(&newPhoto).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating photo"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Photo created successfully"})
// }

// func GetPhotos(c *gin.Context) {
// 	var photos []models.Photo

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Find(&photos).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching photos"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, photos)
// }

// func UpdatePhoto(c *gin.Context) {
// 	photoID := c.Param("photoId")
// 	var updatedPhoto models.Photo

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Where("id = ?", photoID).First(&updatedPhoto).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.Save(&updatedPhoto).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating photo"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
// }

// func DeletePhoto(c *gin.Context) {
// 	photoID := c.Param("photoId")
// 	var photo models.Photo

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Where("id = ?", photoID).First(&photo).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
// 		return
// 	}

// 	if err := db.Delete(&photo).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting photo"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
// }

// func CreatePhoto(c *gin.Context) {
// 	var newPhoto models.Photo
// 	if err := c.ShouldBindJSON(&newPhoto); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := database.GetDB()
// 	defer db.Close()

// 	if err := db.Create(&newPhoto).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Photo created successfully"})
// }

// func GetPhotos(c *gin.Context) {
// 	db := database.GetDB()
// 	defer db.Close()

// 	var photos []models.Photo
// 	if err := db.Find(&photos).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"photos": photos})
// }

// func UpdatePhoto(c *gin.Context) {
// 	photoID := c.Param("photoId")
// 	var updatedPhoto models.Photo

// 	db := database.GetDB()
// 	defer db.Close()

// 	if err := db.Where("id = ?", photoID).First(&updatedPhoto).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.Save(&updatedPhoto).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
// }

// func DeletePhoto(c *gin.Context) {
// 	photoID := c.Param("photoId")

// 	db := database.GetDB()
// 	defer db.Close()

// 	var photo models.Photo
// 	if err := db.Where("id = ?", photoID).First(&photo).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
// 		return
// 	}

// 	if err := db.Delete(&photo).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
// }

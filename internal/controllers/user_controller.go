package controllers

import (
	"net/http"
	"task-5-pbi-btpns-deianearra/internal/config/database"
	"task-5-pbi-btpns-deianearra/internal/helpers"
	"task-5-pbi-btpns-deianearra/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	newUser.Password = string(hashedPassword)

	db := database.DB
	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.DB
	if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("userId")
	var updatedUser models.User

	db := database.DB
	if err := db.Where("id = ?", userID).First(&updatedUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&updatedUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("userId")
	var user models.User

	db := database.DB
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// func RegisterUser(c *gin.Context) {
// 	var newUser models.User
// 	if err := c.ShouldBindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
// 		return
// 	}

// 	newUser.Password = string(hashedPassword)

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Create(&newUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
// }

// func LoginUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	// Generate JWT token
// 	token, err := helpers.GenerateJWT(user.ID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

// func UpdateUser(c *gin.Context) {
// 	userID := c.Param("userId")
// 	var updatedUser models.User

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Where("id = ?", userID).First(&updatedUser).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&updatedUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.Save(&updatedUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
// }

// func DeleteUser(c *gin.Context) {
// 	userID := c.Param("userId")
// 	var user models.User

// 	db, err := database.Connect()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
// 		return
// 	}
// 	defer db.Close()

// 	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := db.Delete(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }

// func createJWTToken(userID uint) (string, error) {
// 	claims := jwt.MapClaims{
// 		"user_id": userID,
// 		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	secretKey := []byte("your-secret-key") // Replace with your actual secret key
// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
// func RegisterUser(c *gin.Context) {
// 	var newUser models.User
// 	if err := c.ShouldBindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := database.GetDB()
// 	defer db.Close()

// 	// Check if email already exists
// 	var existingUser models.User
// 	if db.Where("email = ?", newUser.Email).First(&existingUser).Error == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
// 		return
// 	}

// 	// Hash the password
// 	newUser.Password = helpers.HashPassword(newUser.Password)

// 	if err := db.Create(&newUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
// }

// func LoginUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := database.GetDB()
// 	defer db.Close()

// 	var existingUser models.User
// 	if db.Where("email = ?", user.Email).First(&existingUser).Error != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	// Validate password
// 	if !helpers.CheckPassword(existingUser.Password, user.Password) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	token, err := helpers.GenerateToken(existingUser.ID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

// func UpdateUser(c *gin.Context) {
// 	userID := c.Param("userId")
// 	var updatedUser models.User

// 	db := database.GetDB()
// 	defer db.Close()

// 	if err := db.Where("id = ?", userID).First(&updatedUser).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&updatedUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.Save(&updatedUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
// }

// func DeleteUser(c *gin.Context) {
// 	userID := c.Param("userId")

// 	db := database.GetDB()
// 	defer db.Close()

// 	var user models.User
// 	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := db.Delete(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }

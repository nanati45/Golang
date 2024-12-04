package controller

import (
	"os"
	"task_management_mongoDB/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Global variable to store users in-memory
var users = make(map[string]*models.User)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
	  c.JSON(400, gin.H{"error": "Invalid request payload"})
	  return
	}
  

	// User registration logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
	c.JSON(500, gin.H{"error": "Internal server error"})
	return
	}

	user.Password = string(hashedPassword)
	users[user.Email] = &user

	c.JSON(200, gin.H{"message": "User registered successfully"})
  }


func LoginUser(c *gin.Context) {
	
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
	  c.JSON(400, gin.H{"error": "Invalid request payload"})
	  return
	}
  
	// Global variable to store the JWT secret
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))


	// User login logic
	existingUser, ok := users[user.Email]
	if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
	c.JSON(401, gin.H{"error": "Invalid email or password"})
	return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"user_id": existingUser.ID,
	"email":   existingUser.Email,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return

		}  
		c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})}
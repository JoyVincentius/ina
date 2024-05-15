package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "ina-gin-crud/config"
    "ina-gin-crud/models"
    "time"
)

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
    var users []models.User
    if err := config.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    userID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    var user models.User
    if err := config.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    userID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    var user models.User
    if err := config.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.UpdatedAt = time.Now()
    if err := config.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    userID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    if err := config.DB.Delete(&models.User{}, userID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

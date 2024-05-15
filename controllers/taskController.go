package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "ina-gin-crud/config"
    "ina-gin-crud/models"
    "time"
)

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()
    if err := config.DB.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
    var tasks []models.Task
    if err := config.DB.Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
    id := c.Param("id")
    taskID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    var task models.Task
    if err := config.DB.First(&task, taskID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    taskID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    var task models.Task
    if err := config.DB.First(&task, taskID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.UpdatedAt = time.Now()
    if err := config.DB.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    taskID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    if err := config.DB.Delete(&models.Task{}, taskID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

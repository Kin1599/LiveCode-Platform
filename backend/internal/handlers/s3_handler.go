package handlers

import (
	"fmt"
	"livecode/internal/services/filestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

var s3Client *filestorage.S3Client

func InitS3Client(client *filestorage.S3Client) {
	s3Client = client
}

// UploadProject godoc
// @Summary Загрузка проекта
// @Description Загрузка проекта
// @Tags s3
// @Accept json
// @Produce json
// @Param project_id formData string true "ID проекта"
// @Param project_structure formData string true "Структура проекта"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/uploadProject [post]
func UploadProject(c *gin.Context) {
	fmt.Println(c.Request.PostForm)
	projectID := c.PostForm("project_id")
	projectStructure := c.PostForm("project_structure")

	if err := s3Client.UploadProject(projectID, []byte(projectStructure)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

// DownloadProject godoc
// @Summary Скачивание проекта
// @Description Скачивание проекта
// @Tags s3
// @Accept json
// @Produce json
// @Param project_id query string true "ID проекта"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/downloadProject [get]
func DownloadProject(c *gin.Context) {
	projectID := c.Query("project_id")

	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Project ID"})
		return
	}

	projectStructure, err := s3Client.DownloadProject(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project_structure": string(projectStructure)})
}

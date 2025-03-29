package handlers

import (
	"livecode/internal/services/filestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

var s3Client *filestorage.S3Client

func InitS3Client(client *filestorage.S3Client) {
	s3Client = client
}

type UploadProjectRequest struct {
	ProjectID        string `form:"project_id" binding:"required"`
	ProjectStructure string `form:"project_structure" binding:"required"`
}

// UploadProject godoc
// @Summary Загрузка проекта
// @Description Загрузка проекта
// @Tags s3
// @Accept multipart/form-data
// @Produce json
// @Param project_id formData string true "ID проекта"
// @Param project_structure formData string true "Структура проекта"
// @Success 200 {object} UploadProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/uploadProject [post]
func UploadProject(c *gin.Context) {
	var req UploadProjectRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request parameters"})
		return
	}

	if err := s3Client.UploadProject(req.ProjectID, []byte(req.ProjectStructure)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UploadProjectResponse{Message: "File uploaded successfully"})
}

type DownloadProjectRequest struct {
	ProjectID string `form:"project_id" binding:"required"`
}

// DownloadProject godoc
// @Summary Скачивание проекта
// @Description Скачивание проекта
// @Tags s3
// @Accept json
// @Produce json
// @Param project_id query string true "ID проекта"
// @Success 200 {object} DownloadProjectResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/downloadProject [get]
func DownloadProject(c *gin.Context) {
	var req DownloadProjectRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Missing Project ID"})
		return
	}

	projectStructure, err := s3Client.DownloadProject(req.ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, DownloadProjectResponse{ProjectStructure: string(projectStructure)})
}

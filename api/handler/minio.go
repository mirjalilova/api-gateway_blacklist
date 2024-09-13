package handler

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/mirjalilova/api-gateway_blacklist/pkg/helper"
	"golang.org/x/exp/slog"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Router /file-upload [post]
// @Success 200 {object} string
func (h *HandlerStruct) UploadFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		slog.Error("Error binding request body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ext := filepath.Ext(file.File.Filename)
	if ext != ".pdf" {
		slog.Error("Invalid file format", file.File.Filename)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Only PDF allowed."})
		return
	}

	uploadDir := "./internal/media/"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			slog.Error("Error creating media directory", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating media directory"})
			return
		}
	}

	uploadPath := filepath.Join(uploadDir, file.File.Filename)

	if err := c.SaveUploadedFile(file.File, uploadPath); err != nil {
		slog.Error("Error saving file", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	_, err = h.Minio.Upload(file.File.Filename, uploadPath)
	if err != nil {
		slog.Error("Error uploading file to MinIO", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	filePathTXT, err := helper.ConvertationToTXT(uploadPath)
	if err != nil {
		slog.Error("Error converting PDF to text", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert PDF to text"})
		return
	}

	os.Remove(uploadPath)

	name, err := helper.GetName(h.Genai, filePathTXT)
	if err != nil {
		slog.Error("Error extracting name from PDF", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract name from PDF"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "File successfully uploaded",
		"Full Name": name,
	})
}

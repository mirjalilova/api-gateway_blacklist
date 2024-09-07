package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// Bind file from the request
	err := c.ShouldBind(&file)
	if err != nil {
		slog.Error("Error binding request body", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fayl kengaytmasini olish
	ext := filepath.Ext(file.File.Filename)
	if ext != ".pdf" {
		slog.Error("Invalid file format", file.File.Filename)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Only PDF allowed."})
		return
	}

	// Yuklash direktoriyasini yaratish (agar mavjud bo'lmasa)
	uploadDir := "./internal/media"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			slog.Error("Error creating media directory", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating media directory"})
			return
		}
	}


	fmt.Println("::::;;::;", file.File.Filename)
	// Fayl uchun yangi UUID nomini yaratish
	id := uuid.New().String()
	newFilename := id + ext
	uploadPath := filepath.Join(uploadDir, newFilename)

	// Faylni lokal papkaga saqlash
	if err := c.SaveUploadedFile(file.File, uploadPath); err != nil {
		slog.Error("Error saving file", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	// Fayl yuklanganligi haqida javob qaytarish
	c.JSON(http.StatusCreated, gin.H{
		"message": "File successfully uploaded",
		"path":    uploadPath,
	})
}

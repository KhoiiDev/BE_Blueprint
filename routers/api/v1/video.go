package v1

import (
	"be-hoatieu/pkg/upload"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadVideo(c *fiber.Ctx) error {
	// Nhận file từ form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Không có file được gửi lên",
		})
	}

	// Tạo thư mục nếu chưa có
	if upload.CheckVideoSavePath() {
		if err := upload.MkdirVideoSavePath(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Không thể tạo thư mục lưu video",
			})
		}
	}

	// Tạo tên file duy nhất
	filename := time.Now().Format("20060102150405") + "_" + filepath.Base(file.Filename)

	// Lưu file vào thư mục
	savePath := filepath.Join(upload.GetVideoFullPath(), filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Lỗi khi lưu file",
		})
	}

	return c.JSON(fiber.Map{
		"message":        "Upload thành công",
		"video_filename": filename,
		"video_url":      upload.VideoPath + filename,
	})
}

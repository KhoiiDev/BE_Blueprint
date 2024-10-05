package api

import (
	"be-hoatieu/pkg/upload"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UploadFileSingle(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	if file == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "INVALID_PARAMS",
		})
	}
	fileName := upload.GetImageName(file.Filename) // You can rename this function to something more generic
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + fileName
	// imageName := upload.GetImageName(file.Filename)
	// fullPath := upload.GetImageFullPath()
	// savePath := upload.GetImagePath()
	// src := fullPath + imageName

	// Remove CheckImageExt and CheckImageSize if you need to allow other file types
	if err := upload.CheckImage(fullPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "ERROR_UPLOAD_CHECK_IMAGE_FAIL",
		})
	}

	if err := c.SaveFile(file, src); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "ERROR_UPLOAD_SAVE_IMAGE_FAIL",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":         false,
		"msg":           "SUCCESS",
		"file_url":      upload.GetImageFullUrl(fileName),
		"file_save_url": savePath + fileName,
	})
}

// func UploadFileSingle(c *fiber.Ctx) error {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err,
// 		})
// 	}

// 	if file == nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "INVALID_PARAMS",
// 		})
// 	}

// 	fileName := upload.GetImageName(file.Filename) // You can rename this function to something more generic
// 	fullPath := upload.GetImageFullPath()
// 	savePath := upload.GetImagePath()
// 	src := fullPath + fileName

// 	if !upload.CheckFileExt(fileName) || !upload.CheckImageSize(file) {
// 		fmt.Println(upload.CheckFileExt(fileName))
// 		fmt.Println(upload.CheckImageSize(file))
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "ERROR_UPLOAD_CHECK_FILE_FORMAT",
// 		})
// 	}

// 	err = upload.CheckImage(fullPath) // Consider renaming this function as well
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "ERROR_UPLOAD_CHECK_FILE_FAIL",
// 		})
// 	}

// 	if err := c.SaveFile(file, src); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "ERROR_UPLOAD_SAVE_FILE_FAIL",
// 		})
// 	}

//		return c.Status(fiber.StatusOK).JSON(fiber.Map{
//			"error":         false,
//			"msg":           "SUCCESS",
//			"file_url":      upload.GetImageFullUrl(fileName),
//			"file_save_url": savePath + fileName,
//		})
//	}
func UploadFileMultiple(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	files := form.File["file"]
	if files == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "INVALID_PARAMS",
		})
	}
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	for _, file := range files {
		fileName := upload.GetImageName(file.Filename)
		src := fullPath + fileName

		if !upload.CheckFileExt(fileName) || !upload.CheckFileSize(file) {
			fmt.Println(upload.CheckFileExt(fileName))
			fmt.Println(upload.CheckFileSize(file))
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"msg":   "ERROR_UPLOAD_CHECK_FILE_FORMAT",
			})
		}

		err = upload.CheckImage(fullPath)
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"msg":   "ERROR_UPLOAD_CHECK_FILE_FAIL",
			})
		}

		if err := c.SaveFile(file, src); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"msg":   "ERROR_UPLOAD_SAVE_FILE_FAIL",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":         false,
		"msg":           "SUCCESS",
		"file_url":      upload.GetImagePath(),
		"file_save_url": savePath,
	})
}

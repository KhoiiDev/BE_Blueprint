package main

import (
	"be-hoatieu/models"
	"be-hoatieu/pkg/setting"
	"be-hoatieu/routers"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	setting.Setup()
	models.Setup() // Hàm này sẽ tự động khởi tạo dòng mặc định cho Switch}
}

func main() {
	r := fiber.New(fiber.Config{
		ReadTimeout:  setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		BodyLimit:    setting.ServerSetting.VideoLimit,
	})

	// Gán db cho package v1

	r.Use(cors.New())

	routers.InitRouter(r)

	r.Listen(":3966")
}

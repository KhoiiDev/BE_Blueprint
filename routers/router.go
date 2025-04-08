package routers

import (
	"be-hoatieu/pkg/upload"
	"be-hoatieu/routers/api"
	v1 "be-hoatieu/routers/api/v1"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitRouter(r *fiber.App) {
	// Global Middleware CORS (áp dụng cho toàn bộ server, bao gồm API routes)
	r.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Có thể ghi rõ: "http://localhost:5173"
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	apiv1 := r.Group("/api/v1")

	// dang ky
	apiv1.Post("/signup", v1.SignUp)
	// dang nhap
	apiv1.Post("/signin", v1.SignIn)

	// Compress, Logger, RequestID cho toàn bộ server
	r.Use(compress.New(compress.ConfigDefault))
	r.Use(logger.New(logger.ConfigDefault))
	r.Use(requestid.New())

	fmt.Println(upload.GetImageFullPath())
	fmt.Println(upload.GetVideoFullPath()) // Log đường dẫn video nếu cần

	// Serve static image files
	imageGroup := r.Group("/api/v1/upload/files/images")
	imageGroup.Static("/", upload.GetImageFullPath(), fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Set("Access-Control-Allow-Origin", "*")
			c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			c.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			return nil
		},
	})

	// Serve static video files
	videoGroup := r.Group("/api/v1/upload/files/videos")
	videoGroup.Static("/", upload.GetVideoFullPath(), fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Set("Access-Control-Allow-Origin", "*")
			c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			c.Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			return nil
		},
	})

	// Các API routes khác

	// Nếu cần giới hạn upload size, uncomment dòng này
	// maxUploadSize := int64(10 * 1024 * 1024) // 10MB
	// r.Use(MaxUploadSize(maxUploadSize))

	apiv1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👌!")
	})

	// Upload Image APIs
	apiv1.Post("/upload", api.UploadFileSingle)
	apiv1.Post("/upload/multiple", api.UploadFileMultiple)

	// Upload Video API
	apiv1.Post("/upload-video", v1.UploadVideo)

	// Home page
	home := apiv1.Group("/home")

	// User
	userRoute := apiv1.Group("/user")
	userRoute.Get("/", v1.GetAllUser_Router)
	userRoute.Get("/:id", v1.GetByIdUser_Router)
	userRoute.Get("/username/:username", v1.GetByUserName_Router)
	userRoute.Put("/:id", v1.PutUsers)
	userRoute.Get("/trangthai/status", v1.GetAllUserTrue_Router)
	userRoute.Put("/pass/:id", v1.PutPasswordUsers)
	userRoute.Put("/image/:id", v1.PutIamgeUsers)
	userRoute.Put("/info/:id", v1.PutDetailUsers)
	// Xóa người dùng theo ID
	userRoute.Delete("/:id", v1.DeleteUser) // Thêm route DELETE tại đây

	// introduction
	home.Get("/introduction", v1.GetIntroduction_Component)
	home.Post("/introduction", v1.CreateIntroduction_Component)
	home.Put("/introduction/:id", v1.UpdateIntroduction_Component)
	home.Delete("/introduction/:id", v1.DeleteIntroduction_Component)

	// Home Carousel
	home.Get("/carousel", v1.GetCarousel_Component)
	home.Post("/carousel", v1.CreateCarousel_Component)
	home.Put("/carousel/:id", v1.UpdateCarousel_Component)
	home.Delete("/carousel/:id", v1.DeleteCarousel_Component)

	// Product Service
	home.Get("/servicelist", v1.GetServiceList_Component)
	home.Post("/servicelist", v1.CreateServiceList_Component)
	home.Put("/servicelist/:id", v1.UpdateServiceList_Component)
	home.Delete("/servicelist/:id", v1.DeleteServiceList_Component)

	//News
	home.Get("/news", v1.GetNews_Component)
	home.Post("/news", v1.CreateNews_Component)
	home.Put("/news/:id", v1.UpdateNews_Component)
	home.Delete("news/:id", v1.DeleteNews_Component)

	// Category Navigator
	hoatieu := apiv1.Group("/hoatieu")
	hoatieu.Get("/", v1.GetAllNavigator_Component)
	hoatieu.Post("/", v1.CreateNavigator_Component)
	hoatieu.Put("/:id", v1.UpdateNavigator_Component)
	hoatieu.Delete("/:id", v1.DeleteNavigator_Component)

	// Category product Price
	dichvu := apiv1.Group("/product-price")
	dichvu.Get("/", v1.GetDichvu_Component)
	dichvu.Post("/", v1.CreateDichvu_Component)
	dichvu.Put("/:id", v1.UpdateDichvu_Component)
	dichvu.Delete("/:id", v1.DeleteDichvu_Component)

	// Ship list
	ship := apiv1.Group("/ship")
	ship.Get("/", v1.GetShip_Component)
	ship.Post("/", v1.CreateShip_Component)
	ship.Put("/:id", v1.UpdateShip_Component)
	ship.Delete("/:id", v1.DeleteShip_Component)

	//Tide Calendar
	tideCalendar := apiv1.Group("/tide-calendar")
	tideCalendar.Get("/", v1.GetTideCalendar_Component)
	tideCalendar.Post("/", v1.CreateTideCalendar_Component)
	tideCalendar.Put("/:id", v1.UpdateTideCalendar_Component)
	tideCalendar.Delete("/:id", v1.DeleteTideCalendar_Component)

	// Category Maneuvering Draft
	maneuveringDraft := apiv1.Group("/maneuvering-draft")
	maneuveringDraft.Get("/", v1.GetManeuveringDraft_Component)
	maneuveringDraft.Post("/", v1.CreateManeuveringDraft_Component)
	maneuveringDraft.Put("/:id", v1.UpdateManeuveringDraft_Component)
	maneuveringDraft.Delete("/:id", v1.DeleteManeuveringDraft_Component)

	// Category Maneuvering Draft
	items := apiv1.Group("/items")
	items.Get("/", v1.GetItems_Component)
	items.Post("/", v1.CreateItems_Component)
	items.Put("/:id", v1.UpdateItems_Component)
	items.Delete("/:id", v1.DeleteItems_Component)

	// Category Maneuvering Draft
	kehoachdantau := apiv1.Group("/kehoachdantau")
	kehoachdantau.Get("/", v1.GetKehoachdantau_Component)
	kehoachdantau.Post("/", v1.CreateKehoachdantau_Component)
	kehoachdantau.Put("/:id", v1.UpdateKehoachdantau_Component)
	kehoachdantau.Delete("/:id", v1.DeleteKehoachdantau_Component)

	// Switch routes
	// Switch Maneuvering Draft
	switchGroup := apiv1.Group("/switch")
	switchGroup.Get("/:id", v1.GetSwitch_Component)
	switchGroup.Post("/", v1.CreateSwitch_Component)
	switchGroup.Put("/:id", v1.UpdateSwitch_Component)
	switchGroup.Delete("/:id", v1.DeleteSwitch_Component)

	// Category Maneuvering Draft
	footer := apiv1.Group("/footer")
	footer.Get("/", v1.GetFooter_Component)
	footer.Post("/", v1.CreateFooter_Component)
	footer.Put("/:id", v1.UpdateFooter_Component)
	footer.Delete("/:id", v1.DeleteFooter_Component)

	// Category Maneuvering Draft
	header := apiv1.Group("/header")
	header.Get("/", v1.GetHeader_Component)
	header.Post("/", v1.CreateHeader_Component)
	header.Put("/:id", v1.UpdateHeader_Component)
	header.Delete("/:id", v1.DeleteHeader_Component)
}

// func MaxUploadSize(maxSize int64) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Thiết lập giá trị tối đa vào context để sử dụng trong handler
// 		c.Locals("maxUploadSize", maxSize)
// 		return c.Next()
// 	}
// }

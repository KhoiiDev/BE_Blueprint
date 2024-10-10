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

	apiv1 := r.Group("/api/v1")

	// // dang ky
	// apiv1.Post("/signup", v1.SignUp)
	// // dang nhap
	// apiv1.Post("/signin", v1.SignIn)

	apiv1.Use(cors.New(cors.ConfigDefault))         // CORS
	apiv1.Use(compress.New(compress.ConfigDefault)) // Compress
	apiv1.Use(logger.New(logger.ConfigDefault))     // Logger
	apiv1.Use(requestid.New())
	fmt.Println(upload.GetImageFullPath())

	apiv1.Static("/upload/files", upload.GetImageFullPath(), fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	// JWT check
	// apiv1.Use(jwtCustom.Protected())
	// apiv1.Use(jwtCustom.Authorization())
	// Route
	apiv1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👌!")
	})

	//Image
	apiv1.Post("/upload", api.UploadFileSingle)
	apiv1.Post("/upload/multiple", api.UploadFileMultiple)

	// Home page
	home := apiv1.Group("/home")
	home.Get("/carousel", v1.GetCarousel_Component)
	home.Get("/introduction", v1.GetIntroduction_Component)

	// Product Service
	home.Get("/dichvu", v1.GetServiceList_Component)
	home.Post("/dichvu", v1.CreateServiceList_Component)
	home.Put("/dichvu/:id", v1.UpdateServiceList_Component)
	home.Delete("/dichvu/:id", v1.DeleteServiceList_Component)

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

	// Category Maneuvering Draft
	maneuveringDraft := apiv1.Group("/maneuvering-draft")
	maneuveringDraft.Get("/", v1.GetManeuveringDraft_Component)

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

	//Stream
	stream := apiv1.Group("/stream")
	stream.Get("/", v1.GetStream_Component)
}

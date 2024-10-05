package authentication

import (
	"be-hoatieu/pkg/setting"
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"

	"be-hoatieu/pkg/utils"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(setting.AppSetting.JwtSecret),
		SuccessHandler: Auth,
		ErrorHandler:   jwtError,
		AuthScheme:     "Bearer",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"code":    400,
				"error":   err.Error(),
				"message": "Missing or malformed JWT",
				"data":    nil,
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{
			"code":    401,
			"errors":  err.Error(),
			"message": "Invalid or expired JWT",
			"data":    nil,
		})
}

// Auth is the authentication middleware
func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")
	if h == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: "Authorization header not found."})
	}

	// for example : Bearer xxx.yyy.zzz
	if !strings.HasPrefix(h, "Bearer") {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: `Authorization header is not of type 'Bearer'.`})
	}

	// split the string into 2 parts : 'Bearer ' and the `xxx.yyy.zzz`
	chunks := strings.Split(h, " ")
	if len(chunks) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: `Authorization header value has too many parts. It must follow the pattern: 'Bearer xx.yy.zz' where xx.yy.zz is a valid JWT token.`})
	}

	claims, err := utils.VerifyToken(chunks[1])

	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: "Token has timed out"})
		default:
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: "Token authentication failed"})
		}
	}

	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: "Token authentication failed"})
	}

	// Convert
	profile := jwt.MapClaims(*claims)

	username, err := base64.StdEncoding.DecodeString(profile["username"].(string))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: "Token authentication failed"})
	}
	role, err := base64.StdEncoding.DecodeString(profile["role"].(string))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Error{Code: 401, Message: "Token authentication failed"})
	}
	c.Locals("USER", username)
	c.Locals("ROLE", role)
	return c.Next()
}

// func Authorization() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		val := string(c.Locals("ROLE").([]byte))
// 		if len(val) == 0 {
// 			return c.Status(fiber.StatusUnauthorized).JSON(
// 				&fiber.Error{Message: "user hasn't logged in yet"},
// 			)
// 		}

// 		obj := c.Request().URI().RequestURI()
// 		act := c.Request().Header.Method()
// 		e := models.Casbin()

// 		// // Casbin enforces policy
// 		ok, err := e.Enforce(string(val), string(obj), string(act))

// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(
// 				&fiber.Error{Message: "error occurred when authorizing user"},
// 			)
// 		}
// 		if !ok {
// 			return c.Status(fiber.StatusForbidden).JSON(
// 				&fiber.Error{Message: "forbidden"},
// 			)
// 		}

// 		return c.Next()
// 	}
// }

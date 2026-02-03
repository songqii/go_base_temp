package templates

import "fmt"

func ApiGo(moduleName string) string {
	return fmt.Sprintf(`package controller

import (
	v1 "%s/pkg/controller/v1"
	"%s/conf"
	"%s/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const RequestIDKey = "request_id"

func StartApi() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 MB
	})

	// Middleware
	app.Use(cors.New())
	app.Use(requestid.New(requestid.Config{
		Header:     fiber.HeaderXRequestID,
		ContextKey: RequestIDKey,
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// API v1
	v1Router := app.Group("/v1")
	registerUserRoutes(v1Router)

	// Start server
	log.Zlog.Infof("Server starting on port %%s", conf.Conf.Socket.Port)
	log.Zlog.Fatal(app.Listen(":" + conf.Conf.Socket.Port))
}

func registerUserRoutes(router fiber.Router) {
	router.Post("/login", v1.Login)
	router.Post("/get_user_info", v1.GetUserInfo)
	router.Post("/update_user_info", v1.UpdateUserInfo)
}
`, moduleName, moduleName, moduleName)
}

func UserControllerGo(moduleName string) string {
	return fmt.Sprintf(`package v1

import (
	"%s/log"
	"%s/pkg/error_code"
	"%s/pkg/request"
	"%s/pkg/service"

	"github.com/gofiber/fiber/v2"
)

const RequestIDKey = "request_id"

func Login(c *fiber.Ctx) error {
	req := &request.LoginRequest{}
	reqID := c.Locals(RequestIDKey).(string)

	if err := c.BodyParser(req); err != nil {
		return error_code.ResponseError(c, error_code.ErrorParams)
	}

	log.Zlog.Infof("Login[%%s] request: %%+v", reqID, req)

	resp, err := service.Login(req)
	if err != nil {
		log.Zlog.Errorf("Login[%%s] error: %%+v", reqID, err)
		return error_code.ResponseError(c, err)
	}

	log.Zlog.Infof("Login[%%s] response: %%+v", reqID, resp)
	return c.Status(fiber.StatusOK).JSON(resp)
}

func GetUserInfo(c *fiber.Ctx) error {
	req := &request.GetUserInfoRequest{}
	reqID := c.Locals(RequestIDKey).(string)

	if err := c.BodyParser(req); err != nil {
		return error_code.ResponseError(c, error_code.ErrorParams)
	}

	log.Zlog.Infof("GetUserInfo[%%s] request: %%+v", reqID, req)

	resp, err := service.GetUserInfo(req)
	if err != nil {
		log.Zlog.Errorf("GetUserInfo[%%s] error: %%+v", reqID, err)
		return error_code.ResponseError(c, err)
	}

	log.Zlog.Infof("GetUserInfo[%%s] response: %%+v", reqID, resp)
	return c.Status(fiber.StatusOK).JSON(resp)
}

func UpdateUserInfo(c *fiber.Ctx) error {
	req := &request.UpdateUserInfoRequest{}
	reqID := c.Locals(RequestIDKey).(string)

	if err := c.BodyParser(req); err != nil {
		return error_code.ResponseError(c, error_code.ErrorParams)
	}

	log.Zlog.Infof("UpdateUserInfo[%%s] request: %%+v", reqID, req)

	resp, err := service.UpdateUserInfo(req)
	if err != nil {
		log.Zlog.Errorf("UpdateUserInfo[%%s] error: %%+v", reqID, err)
		return error_code.ResponseError(c, err)
	}

	log.Zlog.Infof("UpdateUserInfo[%%s] response: %%+v", reqID, resp)
	return c.Status(fiber.StatusOK).JSON(resp)
}
`, moduleName, moduleName, moduleName, moduleName)
}

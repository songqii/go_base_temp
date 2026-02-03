package templates

import "fmt"

func ErrorGo(moduleName string) string {
	return fmt.Sprintf(`package error_code

import (
	"fmt"
	"net/http"

	"%s/log"
	"github.com/gofiber/fiber/v2"
)

var ErrorMap = map[error]ErrorInfo{}

type ErrorInfo struct {
	Code    int    `+"`json:\"code\"`"+`
	Message string `+"`json:\"message\"`"+`
}

var (
	Success              = fmt.Errorf("success")
	ErrorInternal        = fmt.Errorf("server internal error")
	ErrorParams          = fmt.Errorf("params error")
	ErrorDataNotFound    = fmt.Errorf("data not found")
	ErrorTokenExpired    = fmt.Errorf("token expired")
	ErrorTokenInvalid    = fmt.Errorf("token invalid")
	ErrorUnauthorized    = fmt.Errorf("unauthorized")
)

func init() {
	ErrorMap = make(map[error]ErrorInfo)
	ErrorMap[Success] = ErrorInfo{Code: 0, Message: Success.Error()}
	ErrorMap[ErrorInternal] = ErrorInfo{Code: 50000, Message: ErrorInternal.Error()}
	ErrorMap[ErrorParams] = ErrorInfo{Code: 40001, Message: ErrorParams.Error()}
	ErrorMap[ErrorDataNotFound] = ErrorInfo{Code: 40004, Message: ErrorDataNotFound.Error()}
	ErrorMap[ErrorTokenExpired] = ErrorInfo{Code: 40005, Message: ErrorTokenExpired.Error()}
	ErrorMap[ErrorTokenInvalid] = ErrorInfo{Code: 40006, Message: ErrorTokenInvalid.Error()}
	ErrorMap[ErrorUnauthorized] = ErrorInfo{Code: 40100, Message: ErrorUnauthorized.Error()}
}

func ResponseError(c *fiber.Ctx, err error) error {
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(ErrorMap[Success])
	}

	log.Zlog.Errorf("ResponseError: %%+v", err)
	if info, ok := ErrorMap[err]; ok {
		return c.Status(http.StatusOK).JSON(info)
	}

	return c.Status(http.StatusOK).JSON(ErrorMap[ErrorInternal])
}
`, moduleName)
}

func RequestGo() string {
	return `package request

type Base struct {
	DeviceID   string ` + "`json:\"device_id\"`" + `
	OsType     string ` + "`json:\"os_type\"`" + `    // ios/android
	AppVersion string ` + "`json:\"app_version\"`" + `
}

type LoginRequest struct {
	Base  Base   ` + "`json:\"base\"`" + `
	Phone string ` + "`json:\"phone\"`" + `
	Code  string ` + "`json:\"code\"`" + `
}

type GetUserInfoRequest struct {
	Base  Base   ` + "`json:\"base\"`" + `
	Token string ` + "`json:\"token\"`" + `
}

type UpdateUserInfoRequest struct {
	Base     Base   ` + "`json:\"base\"`" + `
	Token    string ` + "`json:\"token\"`" + `
	Nickname string ` + "`json:\"nickname\"`" + `
	Avatar   string ` + "`json:\"avatar\"`" + `
}
`
}

func ResponseGo() string {
	return `package response

type BaseResponse struct {
	Code    int    ` + "`json:\"code\"`" + `
	Message string ` + "`json:\"message\"`" + `
}

type LoginResponse struct {
	BaseResponse
	Token string ` + "`json:\"token\"`" + `
	IsNew bool   ` + "`json:\"is_new\"`" + `
}

type UserInfo struct {
	UID      string ` + "`json:\"uid\"`" + `
	Nickname string ` + "`json:\"nickname\"`" + `
	Avatar   string ` + "`json:\"avatar\"`" + `
	Phone    string ` + "`json:\"phone\"`" + `
}

type GetUserInfoResponse struct {
	BaseResponse
	UserInfo UserInfo ` + "`json:\"user_info\"`" + `
}

type UpdateUserInfoResponse struct {
	BaseResponse
}
`
}

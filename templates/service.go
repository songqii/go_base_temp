package templates

import "fmt"

func ServiceGo(moduleName string) string {
	return fmt.Sprintf(`package service

import (
	"%s/pkg/error_code"
	"%s/pkg/utils"
)

// TokenParse parses token and returns user info
func TokenParse(token string) (*utils.JWTClaims, error) {
	if token == "" {
		return nil, error_code.ErrorTokenInvalid
	}

	claims, err := utils.ParseToken(token)
	if err != nil {
		return nil, error_code.ErrorTokenInvalid
	}

	return claims, nil
}
`, moduleName, moduleName)
}

func UserServiceGo(moduleName string) string {
	return fmt.Sprintf(`package service

import (
	"%s/pkg/error_code"
	"%s/pkg/request"
	"%s/pkg/response"
	"%s/pkg/storage"
	"%s/pkg/utils"

	"github.com/google/uuid"
)

func Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	resp := &response.LoginResponse{}
	resp.Code = 0
	resp.Message = "success"

	// Validate phone and code (simplified here)
	if req.Phone == "" {
		return nil, error_code.ErrorParams
	}

	// Find or create user
	user, err := storage.GetUserByPhone(req.Phone)
	if err != nil {
		return nil, err
	}

	isNew := false
	if user == nil {
		// Create new user
		user = &storage.User{
			UID:      uuid.New().String(),
			Phone:    req.Phone,
			Nickname: "User" + req.Phone[len(req.Phone)-4:],
			Status:   1,
		}
		if err := storage.CreateUser(user); err != nil {
			return nil, err
		}
		isNew = true
	}

	// Generate token
	token, err := utils.GenerateToken(user.UID)
	if err != nil {
		return nil, err
	}

	resp.Token = token
	resp.IsNew = isNew
	return resp, nil
}

func GetUserInfo(req *request.GetUserInfoRequest) (*response.GetUserInfoResponse, error) {
	resp := &response.GetUserInfoResponse{}
	resp.Code = 0
	resp.Message = "success"

	claims, err := TokenParse(req.Token)
	if err != nil {
		return nil, err
	}

	user, err := storage.GetUserByUID(claims.UID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, error_code.ErrorDataNotFound
	}

	resp.UserInfo = response.UserInfo{
		UID:      user.UID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Phone:    maskPhone(user.Phone),
	}

	return resp, nil
}

func UpdateUserInfo(req *request.UpdateUserInfoRequest) (*response.UpdateUserInfoResponse, error) {
	resp := &response.UpdateUserInfoResponse{}
	resp.Code = 0
	resp.Message = "success"

	claims, err := TokenParse(req.Token)
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}

	if len(updates) > 0 {
		if err := storage.UpdateUser(claims.UID, updates); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// maskPhone masks phone number
func maskPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}
`, moduleName, moduleName, moduleName, moduleName, moduleName)
}

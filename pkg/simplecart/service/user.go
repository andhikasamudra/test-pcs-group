package service

import (
	"errors"
	"github.com/andhikasamudra/test-pcs-group/internal/env"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/dto"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/repo"
	"github.com/andhikasamudra/test-pcs-group/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Dependency struct {
	SimpleCartModel repo.Interface
}

type SimpleCartService struct {
	SimpleCartModel repo.Interface
}

func NewService(d Dependency) *SimpleCartService {
	return &SimpleCartService{
		SimpleCartModel: d.SimpleCartModel,
	}
}

func (s *SimpleCartService) CreateUser(ctx *fiber.Ctx, request dto.CreateUserRequest) error {
	hashedPassword := utils.HashAndSalt([]byte(request.Password))
	book := repo.User{
		Email:    request.Email,
		Password: hashedPassword,
	}

	_, err := s.SimpleCartModel.CreateUser(ctx.Context(), book)
	if err != nil {
		return err
	}

	return nil
}

func (s *SimpleCartService) LoginRequest(ctx *fiber.Ctx, request dto.LoginRequest) (*dto.LoginResponse, error) {
	userData, err := s.SimpleCartModel.ReadUser(ctx.Context(), repo.GetUserRequest{
		Email: request.Email,
	})
	if err != nil {
		return nil, err
	}

	isPasswordMatched := utils.ComparePassword(userData.Password, []byte(request.Password))

	if !isPasswordMatched {
		return nil, errors.New("invalid email/password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Email
	claims["guid"] = userData.GUID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	resultToken, err := token.SignedString([]byte(env.JWTSecret()))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: resultToken,
	}, nil
}

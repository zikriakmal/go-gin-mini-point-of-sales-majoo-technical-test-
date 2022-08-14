package services

import (
	"errors"
	"majoo/dto"
	"majoo/helper"
	"majoo/repositories"
	"time"
)

type AuthService interface {
	Login(userName string, password string) (dto.LoginResDto, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(repository repositories.UserRepository) AuthService {
	return &authService{
		userRepo: repository,
	}
}
func (s *authService) Login(userName string, password string) (dto.LoginResDto, error) {
	var loginRes dto.LoginResDto
	user, err := s.userRepo.Login(userName, password)

	if err != nil {
		return loginRes, errors.New("401")
	}
	jwt, err := helper.CreateJwt(user)
	if err != nil {
		return loginRes, err
	}

	return dto.LoginResDto{
		Token:     jwt,
		ExpiredAt: time.Now().Unix(),
	}, nil
}

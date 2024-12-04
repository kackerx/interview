package service

import (
	"github.com/kackerx/interview/common/middleware"
)

type Service struct {
	jwt *middleware.JWT
}

func NewService(jwt *middleware.JWT) *Service {
	return &Service{
		jwt: jwt,
	}
}

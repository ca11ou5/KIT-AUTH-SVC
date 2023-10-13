package service

import (
	"Auth_Service/internal/entity"
	"Auth_Service/internal/pb"
	"Auth_Service/internal/repository"
	"context"
	"net/http"
)

type Server struct {
	Repo repository.Repository
}

func (s *Server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	user := entity.User{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		Name:        req.Name,
		Surname:     req.Surname,
		DateOfBirth: req.DateOfBirth,
	}

	err := s.Repo.CreateUser(&user)
	if err != nil {
		return &pb.SignUpResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.SignUpResponse{
		Status: http.StatusCreated,
		Error:  "",
	}, nil
}

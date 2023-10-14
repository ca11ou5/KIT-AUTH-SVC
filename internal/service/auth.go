package service

import (
	"Auth_Service/internal/client"
	"Auth_Service/internal/entity"
	"Auth_Service/internal/pb"
	"Auth_Service/internal/repository"
	"context"
	"net/http"
)

type Server struct {
	Repo   repository.Repository
	Client client.SmsServiceClient
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

func (s *Server) CheckUser(ctx context.Context, req *pb.CheckUserRequest) (*pb.CheckUserResponse, error) {
	err := s.Repo.CheckPhoneNumber(req.PhoneNumber)
	if err != nil {
		return &pb.CheckUserResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	res, err := s.Client.SendCode(req.PhoneNumber)
	if err != nil {
		return &pb.CheckUserResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}
	if res.Error != "" {
		return &pb.CheckUserResponse{
			Status: res.Status,
			Error:  res.Error,
		}, nil
	}

	return &pb.CheckUserResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (s *Server) VerifyPhone(ctx context.Context, req *pb.VerifyPhoneRequest) (*pb.VerifyPhoneResponse, error) {
	return &pb.VerifyPhoneResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

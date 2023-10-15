package service

import (
	"Auth_Service/internal/client"
	"Auth_Service/internal/entity"
	"Auth_Service/internal/pb"
	"Auth_Service/internal/repository"
	"Auth_Service/internal/utils/bcrypt"
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
		Password:    bcrypt.HashPassword(req.Password),
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

	err = s.Repo.SetPhoneCode(req.PhoneNumber, res.Code)
	if err != nil {
		return &pb.CheckUserResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CheckUserResponse{
		Status: res.Status,
		Error:  "",
	}, nil
}

func (s *Server) VerifyPhone(ctx context.Context, req *pb.VerifyPhoneRequest) (*pb.VerifyPhoneResponse, error) {
	err := s.Repo.CheckCode(req.PhoneNumber, req.Code)
	if err != nil {
		return &pb.VerifyPhoneResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	err = s.Repo.ConfirmUser(req.PhoneNumber)
	if err != nil {
		return &pb.VerifyPhoneResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.VerifyPhoneResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (s *Server) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	hashedPassword, err := s.Repo.GetCredentials(req.PhoneNumber)
	if err != nil {
		return &pb.SignInResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	ok := bcrypt.CheckPassword(hashedPassword, req.Password)
	if !ok {
		return &pb.SignInResponse{
			Status: http.StatusUnauthorized,
			Error:  "wrong credentials",
		}, nil
	}

	return &pb.SignInResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

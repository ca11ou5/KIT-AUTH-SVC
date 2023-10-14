package client

import (
	"Auth_Service/internal/pb"
	"context"
	"google.golang.org/grpc"
	"log"
)

type SmsServiceClient struct {
	Client pb.SmsServiceClient
}

func InitSmsClient(url string) SmsServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to establish connection: " + err.Error())
	}

	return SmsServiceClient{Client: pb.NewSmsServiceClient(cc)}
}

func (s *SmsServiceClient) SendCode(phoneNumber string) (*pb.SendCodeResponse, error) {
	req := &pb.SendCodeRequest{PhoneNumber: phoneNumber}

	return s.Client.SendCode(context.Background(), req)
}

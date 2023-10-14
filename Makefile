proto:
	protoc ./internal/pb/auth.proto --go_out=plugins=grpc:.
	protoc ./internal/pb/sms.proto --go_out=plugins=grpc:.

buildimage:
	docker build -t auth-svc .
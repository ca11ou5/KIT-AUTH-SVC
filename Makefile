proto:
	protoc ./internal/pb/auth.proto --go_out=plugins=grpc:.

buildimage:
	docker build -t auth-svc .
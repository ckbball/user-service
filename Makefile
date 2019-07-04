build:
  protoc -I=$(GOPATH)/src/github.com/ckbball/user-service/proto/user --go_out=plugins=micro:$(GOPATH)/src/github.com/ckbball/user-service/proto/user/ $(GOPATH)/src/github.com/ckbball/user-service/proto/user/user.proto
  docker build -t user-service .

run:
  docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 user-service
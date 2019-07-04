package main

import (
  "log"

  pb "github.com/ckbball/user-service/proto/user"
  "github.com/micro/go-micro"
  _ "github.com/micro/go-plugins/registry/mdns"
  k8s "github.com/micro/kubernetes/go/micro"
)

func main() {

  // Creates db connection and handles closing
  db, err := CreateConnection()
  defer db.Close()

  if err != nil {
    log.Fatalf("Could not connect to DB: %v", err)
  }

  // Automatically migrates user struct into db
  db.AutoMigrate(&pb.User{})

  repo := &UserRepository{db}

  tokenService := &TokenService{repo}

  // Create a new service
  srv := k8s.NewService(
    // this name must match package name given in protobuf definition
    micro.Name("user"),
  )

  // Init will parse the command line flags
  srv.Init()

  // Register handler
  pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

  // Run server
  if err := srv.Run(); err != nil {
    log.Fatal(err)
  }
}

package main

import (
  pb "github.com/ckbball/user-service/proto/user"
  "golang.org/x/net/context"
)

type service struct {
  repo         Repository
  tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
  user, err := srv.repo.Get(req.Id)
  if err != nil {
    return err
  }

  res.User = user
  return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.User, res *pb.Response) error {
  users, err := srv.repo.GetAll()
  if err != nil {
    return err
  }
  res.Users = users
  return nil
}

func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Response) error {
  user, err := srv.repo.GetByEmailAndPassword(req)
  if err != nil {
    return err
  }
  res.Token = "testingabc"
  return nil
}

func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
  if err := srv.repo.Create(req); err != nil {
    return err
  }
  res.User = req
  return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *pb.User, res *pb.Response) error {
  return nil
}

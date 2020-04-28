// shippy-service-user/main.go
package main

import (
	pb "github.com/canhdoan/shippy-service-user/proto/user"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) Create(ctx *context.Context, req *pb.User, res *pb.Response) error {
	if err := s.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (s *service) Get(ctx *context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (s *service) GetAll(ctx *context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (s *service) Auth(ctx *context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}

	res.Token = "user valid"
	return nil
}

func (s *service) ValidateToken(ctx *context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}

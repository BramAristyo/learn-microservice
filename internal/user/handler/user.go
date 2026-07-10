package handler

import (
	userpb "github.com/bramAristyo/learn-microservice/proto/user"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

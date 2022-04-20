package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"users/db"
	"users/userspb"

	"google.golang.org/grpc"
)

type Server struct{}

func main() {
	Serve()
}

func Serve() {
	fmt.Println("Server running on port 50051...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Unable to listen on port 50051: ", err)
	}

	s := grpc.NewServer()
	userspb.RegisterUserServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Unable to Serve: ", err)
	}
}

func (*Server) SignUp(ctx context.Context, req *userspb.SignUpRequest) (*userspb.SignUpResponse, error) {
	firstName := req.GetUser().FirstName
	lastName := req.GetUser().LastName
	email := req.GetUser().Email
	userName := req.GetUser().UserName
	password := req.GetUser().Password

	result, err := db.CreateUser(firstName, lastName, email, userName, password)
	if err != nil {
		log.Fatal("There was an issue signing up: ", err)
	}
	res := userspb.SignUpResponse{
		Result: result,
	}
	return &res, nil
}

func (*Server) Delete(ctx context.Context, req *userspb.UserRequest) (*userspb.UserResponse, error) {
	userName := req.GetUser().UserName
	password := req.GetUser().Password

	result, err := db.DeleteUser(userName, password)
	if err != nil {
		log.Fatal("There was an issue deleting user: ", err)
	}
	res := userspb.UserResponse{
		Result: result,
	}
	return &res, nil
}

func (*Server) Login(ctx context.Context, req *userspb.UserRequest) (*userspb.UserResponse, error) {
	userName := req.GetUser().UserName
	password := req.GetUser().Password

	result, err := db.SignIn(userName, password)
	if err != nil {
		log.Fatal("There was an issue with the log in: ", err)
	}
	res := userspb.UserResponse{
		Result: result,
	}
	return &res, nil
}

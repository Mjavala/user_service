package main

import (
	"context"
	"fmt"
	"log"
	"toggl_clone/User/userpb"

	"google.golang.org/grpc"
)

func main() {

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Fatalf("could not log error: %v", err)
	}
	defer cc.Close()

	c := userpb.NewUserServiceClient(cc)
	fmt.Println("Creating User...")
	user := &userpb.User{
		Id:       "user1",
		Name:     "Mjavala",
		Email:    "mjavala@idaf.com",
		Password: "LWJD*(&H&D@R@D#@#",
	}
	u, err := c.CreateUser(context.Background(), &userpb.CreateUserRequest{User: user})
	if err != nil {
		log.Fatalf("Ok big problem: %v", err)
	}

	fmt.Println("User has been created: %v", u)
}

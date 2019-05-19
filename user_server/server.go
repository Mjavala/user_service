package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"toggl_clone/User/userpb"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"google.golang.org/grpc"
)

type server struct{}

type UserItem struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (*server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	fmt.Println("A new user has been requested...")
	user := req.GetUser() //parse data from client

	data := UserItem{ //maps to proto UserItem
		Id:       user.GetId(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}
	fmt.Println("mapping data to doc...")

	doc := map[string]interface{}{
		"_id":      data.Id,
		"Name":     data.Name,
		"Email":    data.Email,
		"Password": data.Password,
	}

	fmt.Println("connecting to couchDB...")
	client, err := kivik.New("couch", "http://0.0.0.0:5984/")
	if err != nil {
		panic(err)
	}

	db := client.DB(context.TODO(), "my_test")

	id := base64.StdEncoding.EncodeToString([]byte(data.Id))

	var rev string
	row := db.Get(context.TODO(), id)
	rev = row.Rev

	if rev == "" {
		fmt.Println("Document does not exist, creating...")
		rev, err = db.Put(context.TODO(), id, doc)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Document exists, updating...")
		doc["_rev"] = row.Rev
		rev, err = db.Put(context.TODO(), id, doc)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Creating gRPC response...")

	return &userpb.CreateUserResponse{ //returns data with ID
		User: &userpb.User{
			Id:       user.GetId(),
			Name:     user.GetName(),
			Email:    user.GetEmail(),
			Password: user.GetPassword(),
		},
	}, nil
}

func main() {

	// if we crash, get the filename and code number

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("User Service Started")

	fmt.Println("Listening...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)

	}
	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	userpb.RegisterUserServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}() //func called

	//wait for CTRL C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is recieved
	<-ch
	fmt.Println("Stopping the Server...")
	s.Stop()
	fmt.Println("Closing Listener...")
	lis.Close()
}

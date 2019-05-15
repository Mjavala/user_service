package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"toggl_clone/User/userpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct{}

type UserItem struct {
	ID       primitive.ObjectID `bson:"_id_omitempty"`
	Name     string             `bson:"id_name"`
	Email    string             `bson:"id_email"`
	Password string             `bson:"id_password"`
}

func (*server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	fmt.Println("A new user has been requested...")
	user := req.GetUser() //parse data

	data := UserItem{ //maps to proto UserItem
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}

	res, err := collection.InsertOne(context.Background(), data) //send to mongo
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID) //create object ID
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to oid: %v", err),
		)
	}

	return &userpb.CreateUserResponse{ //returns data with ID
		User: &userpb.User{
			Id:       oid.Hex(),
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
	fmt.Println("Connecting to MongoDB")

	//connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creating Database")
	//globally accesible collection
	collection = client.Database("userdb").Collection("Users")
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
	fmt.Println("Closing db...")
	client.Disconnect(context.TODO())

}

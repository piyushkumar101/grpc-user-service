package main

import (
	"context"
	"log"
	pb "user_service_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	user, err := client.GetUserDetails(context.Background(), &pb.UserIDRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", user)

	users, err := client.GetUsersDetails(context.Background(), &pb.UserIDsRequest{Ids: []int32{1}})
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	log.Printf("Users: %v", users)

	searchResults, err := client.SearchUsers(context.Background(), &pb.SearchRequest{City: "LA", Married: true})
	if err != nil {
		log.Fatalf("could not search users: %v", err)
	}
	log.Printf("Search Results: %v", searchResults)
}

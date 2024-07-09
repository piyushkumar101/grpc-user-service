package service

import (
	"context"
	"testing"
	pb "user_service_grpc/proto"
)

func TestUserService(t *testing.T) {
	server := NewUserServiceServer()
	user, err := server.GetUserDetails(context.Background(), &pb.UserIDRequest{Id: 1})
	if err != nil || user.User.Id != 1 {
		t.Errorf("expected user id 1, got %v, error: %v", user, err)
	}

	users, err := server.GetUsersDetails(context.Background(), &pb.UserIDsRequest{Ids: []int32{1}})
	if err != nil || len(users.Users) != 1 || users.Users[0].Id != 1 {
		t.Errorf("expected user id 1, got %v, error: %v", users, err)
	}

	searchResults, err := server.SearchUsers(context.Background(), &pb.SearchRequest{City: "LA"})
	if err != nil || len(searchResults.Users) == 0 || searchResults.Users[0].City != "LA" {
		t.Errorf("expected users from LA, got %v, error: %v", searchResults, err)
	}
}

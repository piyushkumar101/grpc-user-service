package service

import (
	"context"
	"fmt"
	pb "user_service_grpc/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	users []pb.User
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		users: []pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		},
	}
}

func (s *UserServiceServer) GetUserDetails(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.UserResponse{User: &user}, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s *UserServiceServer) GetUsersDetails(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	var users []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				users = append(users, &user)
			}
		}
	}
	return &pb.UsersResponse{Users: users}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	var users []*pb.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(user.Married == req.Married) {
			users = append(users, &user)
		}
	}
	return &pb.UsersResponse{Users: users}, nil
}

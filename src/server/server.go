package server

import (
	"context"
	userpb "github.com/asazorojas/go-grpc-users-api/common/protos/v1/user"
	"github.com/asazorojas/go-grpc-users-api/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userpb.UserServer
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userID := req.UserId

	if userID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Length of `userId` cannot be zero")
	}

	var found = false
	var userMessage *userpb.UserMessage
	for _, u := range data.Users {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		found = true
		break
	}

	if !found {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil

}

func (s *UserServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.Users))
	for i, u := range data.Users {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}

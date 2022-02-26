package server

import "context"

//exemplo de um proto
type (
	UserRequest  struct{}
	UserResponse struct{}
	Service      struct{}
)

func (s *Service) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return &UserResponse{}, nil
}

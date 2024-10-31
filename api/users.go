package api

import (
	"context"
	gen "go-time/generated"
)

func ptr(s string) *string {
	return &s
}

func (s Server) GetUsers(ctx context.Context, request gen.GetUsersRequestObject) (gen.GetUsersResponseObject, error) {
	users := []gen.User{
		{
			Email:    ptr("a"),
			Username: "2",
		},
		{
			Email:    ptr("b"),
			Username: "3",
		},
	}

	return gen.GetUsers200JSONResponse(users), nil
}

func (s Server) PostUsers(ctx context.Context, request gen.PostUsersRequestObject) (gen.PostUsersResponseObject, error) {

	return gen.PostUsers200JSONResponse{
		Email:    ptr("a"),
		Username: "2",
	}, nil
}

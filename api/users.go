package api

import (
	"context"
	gen "go-time/generated"
	"go-time/internal/db"
)

func ptr(s string) *string {
	return &s
}

func (s Server) GetUsers(ctx context.Context, request gen.GetUsersRequestObject) (gen.GetUsersResponseObject, error) {
	var p []db.Product

	//db.GetDB().Find(&p)

	var users []gen.User

	for _, v := range p {
		users = append(users, gen.User{
			Email:    ptr(v.Code),
			Username: v.Code,
		})
	}
	return gen.GetUsers200JSONResponse(users), nil
}

func (s Server) PostUsers(ctx context.Context, request gen.PostUsersRequestObject) (gen.PostUsersResponseObject, error) {

	return gen.PostUsers200JSONResponse{
		Email:    ptr("a"),
		Username: "2",
	}, nil
}

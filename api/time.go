package api

import (
	"context"
	"fmt"
	gen "go-time/generated"
	"time"
)

func (s Server) GetTime(ctx context.Context, request gen.GetTimeRequestObject) (gen.GetTimeResponseObject, error) {

	if request.Params.Timezone == nil {
		return nil, fmt.Errorf("timezone parameter is missing")
	}
	tz := *request.Params.Timezone

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("invalid timezone")
	}

	currentTime := time.Now().In(loc)

	return gen.GetTime200JSONResponse{Time: currentTime}, nil
}

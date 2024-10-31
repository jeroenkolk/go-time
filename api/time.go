package api

import (
	"context"
	"fmt"
	gen "go-time/generated"
	"time"
)

func (s Server) GetTime(ctx context.Context, request gen.GetTimeRequestObject) (gen.GetTimeResponseObject, error) {

	tz := *request.Params.Timezone

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("invalid timezone")
	}

	currentTime := time.Now().In(loc)

	return gen.GetTime200JSONResponse{Time: currentTime}, nil
}

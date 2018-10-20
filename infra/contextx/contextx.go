package contextx

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/viriyahendarta/butler-core/infra/constant"
	"github.com/viriyahendarta/butler-core/infra/errorx"
)

type key string

const (
	StartTime key = constant.APPNAME + "_start_time"
	UserID    key = constant.APPNAME + "_user_id"
)

func AppendStartTime(ctx context.Context) context.Context {
	return context.WithValue(ctx, StartTime, time.Now())
}

func GetElapsedTime(ctx context.Context) string {
	if ctx != nil {
		if start, ok := ctx.Value(StartTime).(time.Time); ok {
			elapsed := time.Since(start).Seconds() * 1000
			return fmt.Sprintf("%.2fms", elapsed)
		}
	}
	return "-1ms"
}

func AppendUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserID, userID)
}

func GetUserID(ctx context.Context) (int64, error) {
	sUserID, ok := ctx.Value(UserID).(string)
	if !ok {
		return 0, errorx.New(ctx, errorx.CodeBadRequest, "User ID is empty", nil)
	}
	if sUserID == "" {
		return 0, errorx.New(ctx, errorx.CodeBadRequest, "User ID is empty", nil)
	}

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		return 0, errorx.New(ctx, errorx.CodeParsing, fmt.Sprintf("Failed to parse user id: %s", sUserID), err)
	}
	return userID, nil
}

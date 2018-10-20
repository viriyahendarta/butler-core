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
	AuthID    key = constant.APPNAME + "_auth_id"
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

func AppendAuthID(ctx context.Context, authID string) context.Context {
	return context.WithValue(ctx, AuthID, authID)
}

func GetAuthID(ctx context.Context) (int64, error) {
	sAuthID, ok := ctx.Value(AuthID).(string)
	if !ok {
		return 0, errorx.New(ctx, errorx.CodeBadRequest, "Auth ID is empty", nil)
	}
	if sAuthID == "" {
		return 0, errorx.New(ctx, errorx.CodeBadRequest, "Auth ID is empty", nil)
	}

	authID, err := strconv.ParseInt(sAuthID, 10, 64)
	if err != nil {
		return 0, errorx.New(ctx, errorx.CodeParsing, fmt.Sprintf("Failed to parse Auth ID: %s", sAuthID), err)
	}
	return authID, nil
}

package context

import (
	"context"
	"fmt"
	"time"

	"github.com/viriyahendarta/butler-core/infra/constant"
)

const (
	StartTime string = constant.APPNAME + "_start_time"
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

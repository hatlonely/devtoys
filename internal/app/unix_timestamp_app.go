package app

import (
	"context"
	"math"
	"time"

	"github.com/mergestat/timediff"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type UnixTimestampApp struct {
	ctx context.Context
}

func NewUnixTimestampApp() *UnixTimestampApp {
	return &UnixTimestampApp{}
}

func (a *UnixTimestampApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type TimeInfo struct {
	Timestamp     int64
	TimestampMill int64
	LocalTime     string
	UTCTime       string
	Relative      string
}

func (a *UnixTimestampApp) AnalystTimeInfo(timestr string) (*TimeInfo, error) {
	t, err := strToTimeE(timestr)
	if err != nil {
		return nil, errors.Wrap(err, "cast.ToTimeE failed")
	}

	return &TimeInfo{
		Timestamp:     t.Unix(),
		TimestampMill: t.UnixMilli(),
		LocalTime:     t.Local().Format(time.RFC3339),
		UTCTime:       t.UTC().Format(time.RFC3339),
		Relative:      timediff.TimeDiff(t),
	}, nil
}

func strToTimeE(timestr string) (time.Time, error) {
	if timestr == "" {
		return time.Now(), nil
	}

	i, err := cast.ToInt64E(timestr)
	if err == nil {
		if i < 10000000000 {
			return time.Unix(i, 0), nil
		} else if i < 1000000000000000 {
			return time.Unix(0, i*int64(time.Millisecond)), nil
		} else if i < 1000000000000000000 {
			return time.Unix(0, i*int64(time.Microsecond)), nil
		} else {
			return time.Unix(0, i), nil
		}
	}

	f, err := cast.ToFloat64E(timestr)
	if err == nil {
		i, frac := math.Modf(f)
		if i < 10000000000 {
			return time.Unix(int64(i), int64(frac*1000000000)), nil
		} else if i < 1000000000000000 {
			return time.Unix(0, int64(i*1000000+frac*1000)), nil
		} else if i < 1000000000000000000 {
			return time.Unix(0, int64(i*1000+frac*1000000)), nil
		} else {
			return time.Unix(0, int64(i*frac)), nil
		}
	}

	return cast.ToTimeE(timestr)
}

package devtoys

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

type UnixTimestampReq struct {
	Time string
}

type UnixTimestampRes struct {
	Timestamp     int64
	TimestampMill int64
	LocalTime     string
	UTCTime       string
	Relative      string
	ANSIC         string
	UnixDate      string
	RubyDate      string
	RFC822        string
	RFC822Z       string
	RFC850        string
	RFC1123       string
	RFC1123Z      string
	RFC3339       string
	RFC3339Nano   string
	Kitchen       string
	Stamp         string
}

func (a *UnixTimestampApp) UnixTimestamp(req *UnixTimestampReq) (*UnixTimestampRes, error) {
	t, err := strToTimeE(req.Time)
	if err != nil {
		return nil, errors.Wrap(err, "cast.ToTimeE failed")
	}

	return &UnixTimestampRes{
		Timestamp:     t.Unix(),
		TimestampMill: t.UnixMilli(),
		LocalTime:     t.Local().Format(time.RFC3339),
		UTCTime:       t.UTC().Format(time.RFC3339),
		Relative:      timediff.TimeDiff(t),
		ANSIC:         t.Format(time.ANSIC),
		UnixDate:      t.Format(time.UnixDate),
		RubyDate:      t.Format(time.RubyDate),
		RFC822:        t.Format(time.RFC822),
		RFC822Z:       t.Format(time.RFC822Z),
		RFC850:        t.Format(time.RFC850),
		RFC1123:       t.Format(time.RFC1123),
		RFC1123Z:      t.Format(time.RFC1123Z),
		RFC3339:       t.Format(time.RFC3339),
		RFC3339Nano:   t.Format(time.RFC3339Nano),
		Kitchen:       t.Format(time.Kitchen),
		Stamp:         t.Format(time.Stamp),
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

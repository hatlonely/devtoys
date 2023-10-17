package devtoys

import (
	"context"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnixTimestampApp_AnalystTimeInfo(t *testing.T) {
	a := &UnixTimestampApp{
		ctx: context.Background(),
	}

	Convey("Test UnixTimestampApp.AnalystTimeInfo", t, func() {
		patches := gomonkey.ApplyFunc(time.Now, func() time.Time {
			return time.Unix(1629264000, 0).Add(48 * time.Hour)
		})
		defer patches.Reset()

		tests := []struct {
			name string
			args *UnixTimestampReq
			want *UnixTimestampRes
		}{
			{
				name: "timestamp",
				args: &UnixTimestampReq{
					Time: "1629264000",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "timestamp with milliseconds",
				args: &UnixTimestampReq{
					Time: "1629264000000",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "timestamp with microseconds",
				args: &UnixTimestampReq{
					Time: "1629264000000000",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "timestamp with nanoseconds",
				args: &UnixTimestampReq{
					Time: "1629264000000000000",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339",
				args: &UnixTimestampReq{
					Time: "2021-08-18T05:20:00Z",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339 with milliseconds",
				args: &UnixTimestampReq{
					Time: "2021-08-18T05:20:00.000Z",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339 with timezone",
				args: &UnixTimestampReq{
					Time: "2021-08-18T13:20:00+08:00",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339 with timezone and milliseconds",
				args: &UnixTimestampReq{
					Time: "2021-08-18T13:20:00.000+08:00",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "floating-point timestamp",
				args: &UnixTimestampReq{
					Time: "1629264000.123456",
				},
				want: &UnixTimestampRes{
					Timestamp:     1629264000,
					TimestampMill: 1629264000123,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				got, err := a.UnixTimestamp(tt.args)
				So(err, ShouldBeNil)
				So(got.Timestamp, ShouldEqual, tt.want.Timestamp)
				So(got.TimestampMill, ShouldEqual, tt.want.TimestampMill)
				So(got.LocalTime, ShouldEqual, tt.want.LocalTime)
				So(got.UTCTime, ShouldEqual, tt.want.UTCTime)
				So(got.Relative, ShouldEqual, tt.want.Relative)
			})
		}
	})
}

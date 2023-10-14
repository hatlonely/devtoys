package app

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

		type args struct {
			timestr string
		}
		tests := []struct {
			name string
			args args
			want *TimeInfo
		}{
			{
				name: "timestamp",
				args: args{
					timestr: "1629264000",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "timestamp with milliseconds",
				args: args{
					timestr: "1629264000000",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "timestamp with microseconds",
				args: args{
					timestr: "1629264000000000",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "timestamp with nanoseconds",
				args: args{
					timestr: "1629264000000000000",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339",
				args: args{
					timestr: "2021-08-18T05:20:00Z",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339 with milliseconds",
				args: args{
					timestr: "2021-08-18T05:20:00.000Z",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339 with timezone",
				args: args{
					timestr: "2021-08-18T13:20:00+08:00",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "RFC3339 with timezone and milliseconds",
				args: args{
					timestr: "2021-08-18T13:20:00.000+08:00",
				},
				want: &TimeInfo{
					Timestamp:     1629264000,
					TimestampMill: 1629264000000,
					LocalTime:     "2021-08-18T13:20:00+08:00",
					UTCTime:       "2021-08-18T05:20:00Z",
					Relative:      "2 days ago",
				},
			},
			{
				name: "floating-point timestamp",
				args: args{
					timestr: "1629264000.123456",
				},
				want: &TimeInfo{
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

				got, err := a.AnalystTimeInfo(tt.args.timestr)
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

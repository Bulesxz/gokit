package gokit

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToDateString(t *testing.T) {
	Convey("ToDateString == 20201119", t, func() {
		ts := "20201119"
		now, _ := time.Parse(TIME_YYYYMMDD, ts)
		So(ToDateString(now), ShouldEqual, ts)
	})

	Convey("ToDateString == 20201119", t, func() {
		ts := "20201119"
		now, _ := time.Parse(TIME_YYYY_MM_DD, ts)
		So(ToDateString(now), ShouldNotEqual, ts)
	})
}

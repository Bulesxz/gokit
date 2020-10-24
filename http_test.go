package gokit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseURLQuery(t *testing.T) {
	s := "https://www.baidu.com?k=xxxxxxxx&k2=xxx"
	Convey("ParseURLQuery 获取参数k", t, func() {
		v, err := ParseURLQuery(s, "k")
		So(v, ShouldEqual, "xxxxxxxx")
		So(err, ShouldBeNil)
	})

	Convey("ParseURLQuery 获取参数k3 为空", t, func() {
		v, err := ParseURLQuery(s, "k3")
		So(v, ShouldEqual, "")
		So(err, ShouldBeNil)
	})
	return
}

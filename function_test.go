package gokit

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInarray(t *testing.T) {
	Convey("1 in [1,2,3]", t, func() {
		item := 1
		intArray := []int{1, 2, 3}
		So(InArray(item, intArray), ShouldBeTrue)
	})
	Convey("4 not in [1,2,3]", t, func() {
		item := 4
		intArray := []int{1, 2, 3}
		So(InArray(item, intArray), ShouldBeFalse)
	})
	Convey(`1 panic in ["1","2","3"]`, t, func() {
		item := 1
		intArray := []string{"1", "2", "3"}
		So(func() { InArray(item, intArray) }, ShouldPanic)
	})
}

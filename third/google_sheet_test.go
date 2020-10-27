package third

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewService(t *testing.T) {
	Convey("NewService is nil", t, func() {
		srv, err := NewService("", "", "", "", []string{})
		So(srv, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}

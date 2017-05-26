package modular_cloud_test

import (
	_ "fmt"
	"testing"

	"github.com/pearish/modular_cloud/common"
	"github.com/pearish/modular_cloud/rain"
	"github.com/pearish/modular_cloud/thunder"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCommon(t *testing.T) {
	c := common.NewCloud("first", "second")
	Convey("Given a common cloud", t, func() {
		Convey("The result of A() should be 'first'", func() {
			So(c.A(), ShouldEqual, "first")
		})
		Convey("The result of B() should be 'second'", func() {
			So(c.B(), ShouldEqual, "second")
		})
		Convey("The result of C() should be 'first & second'", func() {
			So(c.C(), ShouldEqual, "first & second")
		})
	})
}

func TestRainCloud(t *testing.T) {
	c := rain.NewRainCloud("first", "second")
	Convey("Given a rain cloud", t, func() {
		Convey("The result of A() should be 'first'", func() {
			So(c.A(), ShouldEqual, "first")
		})
		Convey("The result of B() should be 'second'", func() {
			So(c.B(), ShouldEqual, "second")
		})
		Convey("The result of C() should be 'pitter, patter, ....'", func() {
			So(c.C(), ShouldEqual, "pitter, patter, ....")
		})
	})
}

func TestThunderCloud(t *testing.T) {
	c := thunder.NewThunderCloud("first", "second")
	Convey("Given a thunder cloud", t, func() {
		Convey("The result of A() should be 'first'", func() {
			So(c.A(), ShouldEqual, "first")
		})
		Convey("The result of B() should be 'stronger than a common cloud'", func() {
			So(c.B(), ShouldEqual, "stronger than a common cloud")
		})
		Convey("The result of C() should be 'CRRRRAAAKKK!'", func() {
			So(c.C(), ShouldEqual, "CRRRRAAAKKK!")
		})
	})
	Convey("Treating a thunder cloud as a common cloud", t, func() {
		if cc, ok := c.(common.Cloud); ok {
			Convey("Calling B() directly should be 'stronger than a common cloud'", func() {
				So(cc.B(), ShouldEqual, "stronger than a common cloud")
			})
			Convey("Calling B() from 'base class, i.e. E()' should be 'second'", func() {
				So(cc.E(), ShouldEqual, "second")
			})
			Convey("Calling C() indirectly from base class, i.e. D() should be 'CRRRRAAAKKK!'", func() {
				So(cc.D(), ShouldEqual, "CRRRRAAAKKK!")
			})
		}
	})
}

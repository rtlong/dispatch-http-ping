package main

import (
	// "fmt"

	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEcho(t *testing.T) {

	Convey("Given a `GOECHO` environment variable", t, func() {
		os.Setenv("GOECHO", "this is a test")

		Convey("the output of Echo should be identical to the env var", func() {
			So(Echo(), ShouldEqual, "this is a test")
			os.Setenv("GOECHO", "and another")
			So(Echo(), ShouldEqual, "and another")
		})
	})

	Convey("Given an empty `GOECHO` environment variable", t, func() {
		os.Clearenv()

		Convey("Echo should be an empty string", func() {
			So(Echo(), ShouldEqual, "")
		})
	})

}

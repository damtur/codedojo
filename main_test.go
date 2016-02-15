package main

import (
	. "github.com/smartystreets/goconvey/convey"
	// . "github.com/hudl/alyx3/test"
	// match "github.com/jacobsa/oglematchers"
	// "github.com/jacobsa/oglemock"
	"testing"
)

func TestMain(t *testing.T) {

	Convey("Given code dojo", t, func() {
		So(returnTrue(), ShouldEqual, true)
	})

	Convey("Given a empty list", t, func() {
		Convey("when converting to an slice", func() {

			Convey("should return an empty slice", func() {
				list := NewList()
				So(list.Array(), ShouldResemble, []string{})
			})

		})

		Convey("when appending the string 'foo'", func() {
			Convey("should return ['foo'] when converted to an array", func() {
				list := NewList()
				list.Add("foo")
				So(list.Array(), ShouldResemble, []string{"foo"})
			})
		})

		Convey("when appending the string 'foo' and 'fo1'", func() {
			Convey("should return ['foo', 'fo1'] when converted to an array", func() {
				list := NewList()
				list.Add("foo")
				list.Add("fo1")
				So(list.Array(), ShouldResemble, []string{"foo", "fo1"})
			})
		})

	})
}

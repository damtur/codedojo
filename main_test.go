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
				var list *Node
				So(list.Array(), ShouldResemble, []string{})
			})

		})

		Convey("when appending the string 'foo'", func() {
			Convey("should return ['foo'] when converted to an array", func() {
				var list *Node
				list = list.Add("foo")
				So(list.Array(), ShouldResemble, []string{"foo"})
			})
		})

		Convey("when appending the string 'foo' and 'fo1'", func() {
			Convey("should return ['foo', 'fo1'] when converted to an array", func() {
				var list *Node
				list = list.Add("foo").Add("fo1")
				So(list.Array(), ShouldResemble, []string{"fo1", "foo"})
			})
		})

		Convey("when searching for bar it returns nil", func() {
			var list *Node
			node := list.Find("bar")
			So(node, ShouldEqual, nil)
		})

	})

	Convey("Given a list containing 'foo'", t, func() {
		var list *Node
		list = list.Add("foo")

		Convey("when removing 'foo', returns the empty list", func() {
			list = list.Remove("foo")
			So(list.Array(), ShouldResemble, []string{})
		})
	})

	Convey("Given a list containing 'foo' and 'fo1'", t, func() {
		var list *Node
		list = list.Add("foo").Add("fo1")

		Convey("when removing 'foo', returns the list containing fo1", func() {
			list = list.Remove("foo")
			So(list.Array(), ShouldResemble, []string{"fo1"})
		})

		Convey("when removing 'fo1', returns the list containing foo", func() {
			list = list.Remove("fo1")
			So(list.Array(), ShouldResemble, []string{"foo"})
		})

		Convey("when removing 'fo1' and 'foo', returns the empty list", func() {
			list = list.Remove("fo1").Remove("foo")
			So(list.Array(), ShouldResemble, []string{})
		})
	})

	Convey("Given a list containing 'foo' and 'fo1' and fo2", t, func() {
		var list *Node
		list = list.Add("foo").Add("fo1").Add("fo2")

		Convey("when removing 'foo', returns the list containing fo2, fo1", func() {
			list = list.Remove("foo")
			So(list.Array(), ShouldResemble, []string{"fo2", "fo1"})
		})

		Convey("when removing 'fo1', returns the list containing fo2, foo", func() {
			list = list.Remove("fo1")
			So(list.Array(), ShouldResemble, []string{"fo2", "foo"})
		})

		Convey("when removing 'fo2', returns the list containing fo2, foo", func() {
			list = list.Remove("fo2")
			So(list.Array(), ShouldResemble, []string{"fo1", "foo"})
		})

		Convey("when searching for foo it returns Node with value foo", func() {
			node := list.Find("foo")
			So(node.value, ShouldEqual, "foo")
		})

		Convey("when searching for bar it returns nil", func() {
			node := list.Find("bar")
			So(node, ShouldEqual, nil)
		})
	})
}

package convey

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vektra/mockery/mocks"
	m "github.com/stretchr/testify/mock"
)

func TestNewFunc(t *testing.T)	{
	Convey( "given the blank mock" , t, func() {
		mock := new(mocks.Blank)
		mock.On("Create", m.Anything).Return(nil)
		Convey("when create is clalled", func() {
			err := mock.Create("test")
			Convey("the mock doesn't return an error", func() {
				So(err, ShouldBeNil)
			})
			Convey("the mock can be asserted with https://github.com/smartystreets/goconvey", func() {
				var output string
				Convey("That func shouldn't panic", func() {
					So(func() {
						output = m.NumberOfCalls(mock, "Create", 1)
					}, ShouldNotPanic)
				})
				Convey("That the test passed", func() {
					So(output, ShouldBeBlank)
				})
			})
		})
	})

}
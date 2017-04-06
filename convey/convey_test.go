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
			Convey("That the assertion works", func() {
				Convey("the mock doesn't return an error", func() {
					So(err, ShouldBeNil)
				})
				var output string
				Convey("the func shouldn't panic", func() {
					So(func() {
						output = m.NumberOfCalls(mock, "Create", 1)
					}, ShouldNotPanic)
				})
				Convey("the test returned an empty string", func() {
					So(output, ShouldBeBlank)
				})

				Convey("the mock gets a check", func() {
					So(mock, m.NumberOfCalls, "Create", 1)
				})
			})
		})
	})

}
package convey_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vektra/mockery/mocks"
	m "github.com/stretchr/testify/mock"
	"math/rand"
	my_http "github.com/vektra/mockery/mockery/fixtures/http"
)

func TestBlankMock(t *testing.T)	{
	Convey( "Given the blank mock" , t, func() {
		mock := new(mocks.Blank)
		mock.On("Create", m.Anything).Return(nil)
		Convey("When create is called", func() {
			numberOfCalls := rand.Intn(10) + 1
			for i := 0; i < numberOfCalls; i++ {
				mock.Create("test")
			}
			Convey("NumberOfCalls works as expected", func() {
				Convey("We Pass the assertion", func() {
					So(mock, m.NumberOfCalls, "Create", numberOfCalls)
					mock.AssertNumberOfCalls(t, "Create", numberOfCalls)
				})
				Convey("Returns a string if given the wrong number of calls", func() {
					So(m.NumberOfCalls(mock, "Create", numberOfCalls + 1), ShouldNotBeBlank)
				})
			})
		})
	})
}

func TestExampleMock(t *testing.T)	{
	Convey( "Given the example mock" , t, func() {
		mock := new(mocks.Example)
		mock.On("A").Return(nil)
		mock.On("B", m.Anything).Return(my_http.MyStruct{})
		Convey("When only A is called", func() {
			mock.A()
			Convey("A is called", func() {
				So(mock, m.MethodWasCalled, "A")
				mock.AssertCalled(t, "A")
			})

			Convey("B is not called", func() {
				So(mock, m.MethodWasNotCalled, "B")
				mock.AssertNotCalled(t, "B")
			})
		})
		Convey("When only B is called", func() {
			arg := "test"
			mock.B(arg)
			Convey("B Is called", func() {
				So(mock, m.MethodWasCalled, "B", arg)
			})

			Convey("A is not called", func() {
				So(mock, m.MethodWasNotCalled, "A")
			})
		})
		Convey("When both A and B are called", func() {
			mock.A()
			mock.B("test")
			Convey("Expectations Are Met", func() {
				So(mock, m.ExpectationsWereMet)
				mock.AssertExpectations(t)
			})
		})
	})
}
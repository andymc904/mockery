package convey

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/vektra/mockery/mocks"
	m "github.com/stretchr/testify/mock"
)

func TestExampleMock(t *testing.T)	{
	Convey( "Given the example mock" , t, func() {
		mock := new(mocks.Example)
		mock.On("A").Return(nil)
		Convey("When only A is called", func() {
			mock.A()
			Convey("A is called", func() {
				So(mock, m.MethodWasCalled, "A")
			})
		})
	})
}
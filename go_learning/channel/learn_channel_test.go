package channel

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_doTestChanWithoutCache(t *testing.T) {
	Convey("Test_doTestChanWithoutCache", t, func() {

		doTestChanWithoutCache()
		fmt.Println()

	})
}

func Test_doTestCommunicationWithSync(t *testing.T) {
	Convey("Test_doTestCommunicationWithSync", t, func() {
		doTestCommunicationWithSync()
	})
}

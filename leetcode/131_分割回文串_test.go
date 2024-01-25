package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_partition(t *testing.T) {
	convey.Convey("Test_partition", t, func() {
		fmt.Println(partition("aab"))
	})
}

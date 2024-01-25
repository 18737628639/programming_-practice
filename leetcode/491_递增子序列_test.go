package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	convey.Convey("Test_findSubsequences", t, func() {
		fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
	})
}

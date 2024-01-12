package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	convey.Convey("Test_combinationSum2", t, func() {
		fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	})
}

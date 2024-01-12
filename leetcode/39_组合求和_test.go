package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	convey.Convey("Test_combinationSum", t, func() {
		fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	})
}

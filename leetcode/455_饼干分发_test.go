package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	convey.Convey("Test_findContentChildren", t, func() {
		fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	})
}

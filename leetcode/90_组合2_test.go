package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	convey.Convey("Test_doSubsetWithDup", t, func() {
		fmt.Println(subsetsWithDup([]int{2, 1, 2, 1, 3}))
	})
}

func Test_setNonDupSubset(t *testing.T) {
	convey.Convey("Test_doSubsetWithDup", t, func() {
		fmt.Println(setNonDupSubset([]int{1, 2}, [][]int{{1, 2}, {1, 2, 2}}))
	})
}

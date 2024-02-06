package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	convey.Convey("Test_permuteUnique", t, func() {
		fmt.Println(permuteUnique([]int{1, 1, 2}))
	})
}

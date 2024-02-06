package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_permute(t *testing.T) {
	convey.Convey("Test_permute", t, func() {
		fmt.Println(permute([]int{1, 2, 3}))
	})
}

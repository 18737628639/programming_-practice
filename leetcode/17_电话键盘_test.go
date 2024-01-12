package leetcode

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	convey.Convey("letterCombinations", t, func() {
		fmt.Println(letterCombinations("23"))
	})
}

package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	return doRestore(s, 0 , len(s)-1, nil, "")
}

func doRestore(s string, start, end int,ret []string, addr string)[]string {
	if start >= len(s) && strings.Count(addr, ".") == 4{
		ret = append(ret, addr[:len(addr)-1])
		return ret
	}
	for i := start; i<=end; i ++ {
		tmp := s[start:i+1]
		num, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			continue
		}
		if num > 255 {
			return ret
		}
		if num <0 {
			continue
		}
		if isContinueZeroInStart(tmp){
			continue
		}
		if isZeroStartNum(tmp) {
			continue
		}
		addr += tmp
		ret = doRestore(s, i+1, end, ret, fmt.Sprintf("%s.",addr))
		addr = addr[:len(addr) - len(tmp)]
	}
	return ret
}

func isContinueZeroInStart (s string ) bool {
	var cnt int
	for i := 0 ; i < len(s);i ++ {
		if s[i] == '0' {
			cnt ++
		}else {
			break
		}
	}
	if cnt > 1 {
		return true
	}
	return false
}


func isZeroStartNum (s string ) bool {
	var flag bool
	if len(s) >= 2 && s[0] == '0'{
		return true
	}
	return flag
}

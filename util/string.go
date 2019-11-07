package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var (
	gbkRegexp = regexp.MustCompile("[\u4e00-\u9fa5]")
)

// IsEmptyStr returns true if val == nil or val == "" or val == " "
func IsEmptyStr(val string) bool {
	return strings.TrimSpace(val) == ""
}

func SubStr(str string, begin, length int) string {
	rs := []rune(str)
	lth := len(rs)

	if length < 0 {
		length = lth
	}
	if begin < 0 {
		begin = 0
	}

	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	return string(rs[begin:end])
}

// StrArrayContains returns true if array contains the item
func StrArrayContains(arr []string, item string) bool {
	for _, s := range arr {
		if s == item {
			return true
		}
	}
	return false
}

func StrArrayEquals(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for _, s := range arr1 {
		if !StrArrayContains(arr2, s) {
			return false
		}
	}
	return true
}

func GenRandomString(seed string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := []byte(fmt.Sprintf("%s%d", seed, r.Intn(1000000)))
	return fmt.Sprintf("%x", md5.Sum(data))
}

/**
  @comment check if gbk characters in string
*/
func ContainsGBK(str string) bool {

	return gbkRegexp.MatchString(str)
}

func SafeReplaceName(name string) string {

	result := name
	if !ContainsGBK(name) {
		if arr := strings.SplitN(name, " ", 2); len(arr) > 1 && !IsEmptyStr(arr[1]) {
			result = fmt.Sprintf(`%s %s`, arr[0], strings.ToUpper(arr[1][:1]))
		}
	}
	return result
}

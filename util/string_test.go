package util

import "testing"

func TestSafeReplaceName(t *testing.T) {

	arr := map[string]string{
		"Android 龙":           "Android 龙",
		"Dellinger zhang":     "Dellinger Z",
		"马 达":                 "马 达",
		"乡 wa da na":          "乡 wa da na",
		"Mark williams tomas": "Mark W",
		"Luca S salong":       "Luca S",
		"IT Session":          "IT S",
	}

	for k, v := range arr {
		str := SafeReplaceName(k)
		if str != v {
			t.Errorf("actual:%s expect:%s", str, v)
			t.Fail()
		}
	}
}

func TestContainsGBK(t *testing.T) {

	arr := map[string]bool{
		"dksjfksdj你大爷在梦里dssd":                  true,
		"s9ijksdjfsjdf":                            false,
		"明确要求经济上低价覅哦文件日诶我今晚i二进位":   true,
		"238u2348u38282391u19u":                    false,
		"^%&*^*)*^^%$%$%**(":                       false,
		"☎☏✄☪☣☢☠♨« »큐〓㊚㊛囍㊒㊖☑✔☐☒✘㍿☯☰☷♥♠♤❤♂♀★☆☯✡※卍卐": true,
		"乡 wa da na":                               true,
	}

	for k, v := range arr {
		if ContainsGBK(k) != v {
			t.Errorf("actual:%v, expect:%v", k, v)
			t.Fail()
		}
	}
}

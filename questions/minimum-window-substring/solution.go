package advantage_shuffle

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	//s 和 t 由英文字母组成
	need := map[byte]int{}
	window := map[byte]int{}

	for i := range t {
		need[t[i]]++
	}

	valid := 0
	start := -1
	slen := len(s) + 1

	for left, right := 0, -1; left < len(s) && right+1 < len(s); {
		right = right + 1
		c := s[right]
		if v, ok := need[c]; ok {
			window[c]++
			if window[c] == v {
				valid++
			}
		}

		if valid != len(need) {
			continue
		}

		for left <= right && left < len(s) {
			c = s[left]
			if v, ok := need[c]; ok {
				if len(need)-valid == 0 {
					if slen > right-left {
						start = left
						slen = right - left + 1
					}

					if window[c] == v {
						valid--
					}
					window[c]--
				} else {
					break
				}
			}
			left++
		}

	}

	if start < 0 {
		return ""
	}

	return s[start : start+slen]
}

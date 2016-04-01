package kata

import "strings"

func IsAnagram(s1 string, s2 string) bool {
	s1 = strings.Trim(s1, " ");
	s2 = strings.Trim(s2, " ");
	if (len(s1) != len(s2)) {
		return false;
	}
	m := make(map[string]int)

	for _, r := range s1 {
		c := string(r)
		if m[c] > 0 {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	for _, r := range s2 {
		c := string(r)
		if m[c] > 0 {
			m[c]--
		} else {
			m[c] = 1
		}
	}

	for _, v := range m {
		if v > 0 {
			return false;
		}
	}

	return true;


}
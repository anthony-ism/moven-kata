package kata

func HasUniqueCharacters(s string) bool {
	set := make(map[string]bool)
	for _, r := range s {
		c := string(r)
		if set[c] {
			return false
		} else {
			set[c] = true
		}
	}
	return true;
}

func IsAnagram(s1 string, s2 string) bool {
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
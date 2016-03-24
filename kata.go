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

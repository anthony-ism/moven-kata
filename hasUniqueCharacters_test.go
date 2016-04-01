package kata

import ("testing")

func TestHasUniqueCharacters(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"Hello, world", false},
		{"IamDistnc", true},
		{"", true},
	}
	for _, c := range cases {
		got := HasUniqueCharacters(c.in)
		if got != c.want {
			t.Errorf("HasUniqueCharacters(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
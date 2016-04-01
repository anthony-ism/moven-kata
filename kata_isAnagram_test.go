package kata

import ("testing")

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s1 string
		s2 string
		want bool
	}{
		{"Hello, world", "ello, worldH", true},
		{"IamDistnc", "blahblah", false},
		{"", "", true},
		{"1", "1", true},
	}
	for _, c := range cases {
		got := IsAnagram(c.s1, c.s2)
		if got != c.want {
			t.Errorf("IsAnagram(%q, %q) == %q, want %q", c.s1, c.s2, got, c.want)
		}
	}
}
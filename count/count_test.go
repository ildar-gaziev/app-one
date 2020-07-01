package count

import "testing"

func TestCount(t *testing.T) {
	s := "qwertyeege"
	e := 4
	if c := Count(s, 'e'); c != e {
		t.Fatalf("bad count for %s: got %d expected %d", s, c, e)
	}
}

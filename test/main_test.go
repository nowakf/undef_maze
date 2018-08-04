package main

import (
	"math/rand"
	"testing"
	"time"
)

var s segment

func TestLength(t *testing.T) {
	s.child = &segment{}
	s.child.child = &segment{}
	s.child.child.child = &segment{}
	if s.length() != 4 {
		t.Fatalf("expected 3, got %d", s.length())
	}
}

func Test_unit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn
	largeInt := 100000
	for i := 0; i < 1000; i++ {
		v := vector{r(largeInt) - r(largeInt), r(largeInt) - r(largeInt)}

		u := v.unit()

		if (u.x == 0 || u.x == 1 || u.x == -1) && (u.y == 0 || u.y == 1 || u.y == -1) {
			if (v.x < 0 && u.x < 0) || (v.x > 0 && u.x > 0) || (v.x == 0 && u.x == 0) {
				//it works, probably
			} else {
				t.Fatalf("vector V(%d %d) produced %d %d", v.x, v.y, u.x, u.y)
			}
		} else {
			t.Fatalf("vector V(%d %d) produced %d %d", v.x, v.y, u.x, u.y)
		}
	}
}

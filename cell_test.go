package main

import (
	"math/rand"
	"testing"
	"time"
)

var objects = []object{
	tendril, body, wall, food,
}

var shuffle = func() {
	rand.Shuffle(len(objects), func(i, j int) {
		objects[i], objects[j] = objects[j], objects[i]
	})
}

func TestBitHas(t *testing.T) {
	var empty cell
	shuffle()
	for _, o := range objects {
		if empty.Has(o) {
			t.Fatalf("object found in empty")
		}
	}
	for i := 0; i < 1000; i++ {
		var c cell
		shuffle()
		checks := make([]object, 0)
		for _, o := range objects {
			if rand.Int()%2 > 0 {
				c.Write(o)
				checks = append(checks, o)
			}
		}
		for _, o := range checks {
			if !c.Has(o) {
				t.Fatalf("expected obj not found")

			}

		}
	}
}

func TestWrite(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var c cell
	c.Write(tendril | body | wall | food)
	if !c.Has(tendril | body | wall | food) {
		t.Fatalf("expected %x, got %x", (tendril | body | wall | food), c)
	}

	c = 0

	for i := 0; i < 1000; i++ {
		shuffle()
		checks := make([]object, 0)
		for _, object := range objects {
			if rand.Int()%2 > 0 {
				c.Write(object)
			}
		}
		for _, check := range checks {
			if !c.Has(check) {
				t.Fatalf("expected contents %s, got %s", check.String(), c.Enumerate())
			}
		}
		for _, check := range checks {
			c &^= cell(check)
			if c.Has(check) {
				t.Fatalf("expected contents %s, got %s", check.String(), c.Enumerate())
			}
		}
	}

}
func TestRune(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var c cell
		shuffle()
		for _, object := range objects {
			if rand.Int()%2 > 0 {
				c.Write(object)
			}
		}
		switch {
		case c.Has(wall):
			if c.Rune() != block {
				t.Fatalf("expected %d, got %d", block, c.Rune())
			}

		case c.Has(food):
			if c.Rune() != fmark {
				t.Fatalf("expected %d, got %d", fmark, c.Rune())
			}

		case c.Has(body):
			if c.Rune() != crmark {
				t.Fatalf("expected %d, got %d", crmark, c.Rune())
			}

		case c.Has(tendril):
			//return segment(c).Rune()
		case c == 0:
			if c.Rune() != space {
				t.Fatalf("expected %d, got %d", space, c.Rune())
			}

		}
	}
}

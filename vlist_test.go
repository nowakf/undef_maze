package main

import (
	"math/rand"
	"testing"
)

var lis vlist

var empty vlist

const (
	MAX = 10
)

func TestLen(t *testing.T) {
	lis := vlist{}
	expectedLen := 0
	if lis.Len() != expectedLen {
		t.Fatalf("list length is %d, expected %d", lis.Len(), expectedLen)
	}
	lis = empty
	expectedLen = 0
	if lis.Len() != expectedLen {
		t.Fatalf("list length is %d, expected %d", lis.Len(), expectedLen)
	}
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{0, 0})
		lis.PushFront(vector{0, 0})
	}
	expectedLen = (MAX * 2)
	if lis.Len() != expectedLen {
		t.Fatalf("list length is %d, expected %d", lis.Len(), expectedLen)
	}

}

func TestMap(t *testing.T) {
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{0, 0})
		lis.PushFront(vector{0, 0})
	}
	lis.Map(func(v *vector) {
		*v = v.add(vector{1, 1})
	})
	for i := 0; i < lis.Len(); i++ {
		v3 := vector{1, 1}
		v2 := *lis.I(i)
		if v2 != v3 {
			t.Fatalf("%s should be %s", v2.String(), vector{1, 1}.String())

		}

	}
}
func TestI(t *testing.T) {
	lis = empty

	//test normal case
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{MAX + i, 0})
	}
	for i := 0; i < MAX; i++ {
		lis.PushFront(vector{MAX - (i + 1), 0})
	}

	for i := 0; i < lis.Len(); i++ {
		if lis.I(i).x != i {
			t.Fatalf("entry did not match, I(i)=%d i=%d", lis.I(i).x, i)
		}
	}

	//test nil list case
	lis = empty
	for i := 0; i < lis.Len(); i++ {
		if lis.I(i) != nil {
			t.Fatalf("entry was nil, %d ", i)
		}
	}
	i := 2
	if lis.I(i) != nil {
		t.Fatalf("entry should nil, %d ", i)
	}
	//test unbalanced list case
	lis = empty
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{0, 0})
	}
	for i := 0; i < lis.Len(); i++ {
		if lis.I(i) == nil {
			t.Fatalf("entry was nil, %d ", i)
		}
	}
	//test unbalanced list case 2
	lis = empty
	for i := 0; i < MAX; i++ {
		lis.PushFront(vector{0, 0})
	}
	for i := 0; i < lis.Len(); i++ {
		if lis.I(i) == nil {
			t.Fatalf("entry was nil, %d ", i)
		}
	}

}
func TestHas(t *testing.T) {
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{0, 0})
		lis.PushFront(vector{0, 0})
	}
	testInc := rand.Int() % MAX
	*lis.I(testInc) = vector{1, 2}
	if i, ok := lis.Has(vector{1, 2}); !ok || i != testInc {
		t.Fatalf("test failed with %d, %d, %t", i, testInc, ok)
	}
}
func TestPopBack(t *testing.T) {
	lis = empty
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{i, i})
		lis.PushFront(vector{0, i})
	}
	//max := lis.Len()
	for lis.Len() > 0 {
		v1 := lis.I(lis.Len() - 1)
		v2 := lis.PopBack()
		if *v1 != v2 {
			t.Fatalf("unexpected element popped %s, should be %s", v2.String(), v1.String())
		}
	}
	if lis.Len() != 0 {
		t.Fatalf("length should be 0, is %d", lis.Len())
	}
}
func TestPopFront(t *testing.T) {
	lis = empty
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{i, i})
		lis.PushFront(vector{0, i})
	}
	for lis.Len() > 0 {
		v1 := lis.I(0)
		v2 := lis.PopFront()
		if *v1 != v2 {
			t.Fatalf("unexpected element popped %s, should be %s", v2.String(), v1.String())
		}
	}
	if lis.Len() != 0 {
		t.Fatalf("length should be 0, is %d", lis.Len())
	}
}

//skookum
func TestCut(t *testing.T) {
	lis = empty
	for i := 0; i < MAX; i++ {
		lis.PushBack(vector{0, 0})
		lis.PushFront(vector{0, 0})
	}
	y := MAX * 2
	x := 0
	lis.Map(func(v *vector) {
		v.x = x
		v.y = y
		x++
		y--
	})
	for i := 0; i < lis.Len(); i++ {
		a, b := lis.Cut(i)
		for i := 0; i < a.Len(); i++ {
			if a.I(i) != lis.I(i) {
				t.Fatalf("a = %s, lis = %s, in a list with %s", a.I(i).String(), lis.I(i).String(), lis.Print())
			}
		}
		for i := 0; i < b.Len(); i++ {
			if b.I(i) != lis.I(a.Len()+i) {
				t.Fatalf("b = %s, lis = %s, in a list with %s", b.I(i).String(), lis.I(i).String(), lis.Print())
			}
		}
	}

}

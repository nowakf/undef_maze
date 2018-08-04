package main

import "fmt"

//some operations on vector lists
type vlist struct {
	head []vector
	tail []vector
}

func (v *vlist) Print() string {
	s := ""
	for i := 0; i < v.Len(); i++ {
		s += v.I(i).String()
	}
	return s
}

func (v *vlist) Map(f func(v *vector)) {
	for i := 0; i < v.Len(); i++ {
		f(v.I(i))
	}

}
func (v *vlist) I(i int) *vector {
	switch {
	case i < 0:
		panic(fmt.Sprintf("negative value! %d", i))
	case i < len(v.head):
		//we're in the first half
		return &v.head[len(v.head)-(i+1)]
	case i < len(v.head)+len(v.tail):
		//we're in the second half
		return &v.tail[i-len(v.head)]
	default:
		panic(fmt.Sprintf("access out of bounds, %d", i))
	}

}

func (v *vlist) Len() int {
	return len(v.head) + len(v.tail)
}

func (v *vlist) Has(v1 vector) (int, bool) {
	for i := 0; v.Len() > i; i++ {
		if v1 == *v.I(i) {
			return i, true
		}
	}
	return 0, false
}

//cut returns two lists, the second of which includes 'at'
func (v *vlist) Cut(at int) (part1 vlist, part2 vlist) {
	if at >= v.Len() || at < 0 {
		panic(fmt.Sprintf("tried to cut out of bounds, %d", at))
	}

	if at < len(v.head) {
		part1 = vlist{v.head[len(v.head)-at:], nil}
		part2 = vlist{v.head[:len(v.head)-at], v.tail}
		return
	}
	if at >= len(v.head) {
		part1 = vlist{v.head, v.tail[:at-len(v.head)]}
		part2 = vlist{nil, v.tail[at-len(v.head):]}
		return
	}
	panic("...")
}

func (v *vlist) PushFront(v1 vector) {
	v.head = append(v.head, v1)
}
func (v *vlist) PushBack(v1 vector) {
	v.tail = append(v.tail, v1)
}

func (v *vlist) pop(vecs1 *[]vector, vecs2 *[]vector) (v1 vector) {
	switch {
	case len(*vecs1) > 1:
		*vecs1, v1 = (*vecs1)[:len(*vecs1)-1], (*vecs1)[len(*vecs1)-1]
	case len(*vecs1) == 1:
		*vecs1, v1 = make([]vector, 0), (*vecs1)[0]
	case len(*vecs2) > 1:
		*vecs2, v1 = (*vecs2)[1:], (*vecs2)[0]
	case len(*vecs2) == 1:
		*vecs2, v1 = make([]vector, 0), (*vecs2)[0]
	default:
		panic("pop called on empty list")
	}
	return
}
func (v *vlist) PopBack() (v1 vector) {
	v1 = v.pop(&v.tail, &v.head)
	return
}

func (v *vlist) PopFront() (v1 vector) {
	v1 = v.pop(&v.head, &v.tail)
	return
}

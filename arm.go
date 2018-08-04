package main

type arm struct {
	segments vlist
	path     vlist
	state    astate
}
type astate int

const (
	extending astate = iota
	retracting
)

func (a *arm) Print(vs ...vlist) string {
	s := "SEGMENTS:"
	s += a.segments.Print()
	s += "\n PATH:"
	s += a.path.Print()
	return s
}
func (a *arm) assign(p vlist) {
	a.state = retracting
	a.path = p
}

func (a *arm) move(newLoc vector) {
	a.segments.PushFront(newLoc)
	a.path.PushFront(newLoc)
}

func (a *arm) updateLength() {
	switch a.state {
	case extending:
		if a.path.Len() > a.segments.Len() {
			a.segments.PushBack(*a.path.I(a.segments.Len()))
		}
		if a.path.Len() <= 0 {
			a.state = retracting
		}
	case retracting:
		if a.segments.Len() > 0 {
			a.segments.PopBack()
		} else {
			a.state = extending
		}
	}
}

func (a *arm) draw(m *maze) {
	for i := 0; i < a.segments.Len(); i++ {
		m.Write(*a.segments.I(i), tendril)
	}
	//	for i := 0; i < a.path.Len(); i++ {
	//		m.Write(*a.path.I(i), tendril)
	//	}

}

package main

type object uint16

func (o object) String() string {
	switch o {
	//case body:
	//	return "body"
	case wall:
		return "wall"
	case food:
		return "food"
	case tendril:
		return "tendril"
	}
	return "unknown"
}

const (
	tendril object = 1 << iota
	body_w
	body_c
	body_e
	body_r
	food
	wall
	cloud
	numberOfObjects = iota
)

const (
	USIZE        = 16
	all   object = 0xFFFF >> (USIZE - numberOfObjects)
)

type cell uint16

//you can query however many objects using bitwise OR
//returns 0 if there's no such thing here
func (c cell) Some(query object) (contains object) {
	return object(c) & query
}

//you can query however many objects using bitwise OR
//returns true if cell contains all of them
func (c cell) Has(query object) bool {
	return query&^object(c) == 0
}

//returns a string of all the objects in the cell
func (c cell) Enumerate() (content string) {
	for i := 0; i < int(numberOfObjects); i++ {
		var this object
		this = 1 << uint(i)
		if this&object(c) != 0 {
			content += this.String()
			content += "\n"
		}
	}
	return
}

//or write however many objects using bitwise OR
func (c *cell) Write(o object) {
	*c |= cell(o)
}

//or erase however many objects using bitwise OR
func (c *cell) Erase(o object) {
	*c &^= cell(o)
}

const (
	block    = '\u2588'
	space    = ' '
	fmark    = '!'
	testMark = 'X'
	body_cm  = '\u0473'
	body_em  = '\u0473'
	body_rm  = '\u0473'
	body_wm  = '\u0473'
	crmark   = '\u0473'
	arm_NS   = '\u2503'
	arm_EW   = '\u2501'
	arm_SW   = '\u250F'
	arm_SE   = '\u2513'
	arm_NW   = '\u2517'
	arm_NE   = '\u251B'
)

func (c cell) Rune() rune {
	switch {
	case c.Has(wall):
		return block
	case c.Has(food):
		return fmark
	//case c.Has(body):
	//	return crmark
	case c.Has(tendril):
		return '.'
	case c == 0:
		return space
	}
	return rune(0)
}

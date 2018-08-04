package main

import (
	"errors"
	//"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var mazeSymbols = []rune{
	block,
	crmark,
	arm_EW,
	arm_NE,
	arm_NS,
	arm_NW,
	arm_SE,
	arm_SW,
}

var testmaze = []string{
	"###########################################################################################",
	"#      #                        #                  #                        #             #",
	"#  # #### ############ ###### ###### ###### ## # #### ############ ###### ###### ###### ###",
	"#  # #                    !                    # #                                        #",
	"#  # # ############### ###### ###### ###### ## # # ############### ###### ###### ###### ###",
	"#  # # #             # #    # #    # # !  # #  # # #   !         # #    # #    # #  ! # # #",
	"#  # # #             # #    # #    # #    # #  # # #             # #    # #    # #    # # #",
	"#  # # #             # #    # #    # #    # #  # # #             # #    # #    # #    # # #",
	"#  # # #             # #    # #    # #    # #  # # #             # #    # #    # #    # # #",
	"#  # # ########### # # #  # # #  # # #  # # #  # # ########### # # #  # # #  # # #  # # # #",
	"#  # # #           # # #  # # #  # # #  # # #  # # #           # # #  # # #  # # #  # # # #",
	"#       !          #      #      #      #                      #      #      #      #     #",
	"#   X                                                        !               !            #",
	"#            !                           !                                                #",
	"#      #                        #                  #                        #             #",
	"#  # #### ############ ###### ###### ###### ## # #### ############ ###### ###### ###### ###",
	"#  # #           !            !                # #                                        #",
	"#  # # ############### ###### ###### ###### ## # # ############### ###### ###### ###### ###",
	"#  # # #             # #    # #    # #    # #  # # #             # #    # #    # #    # # #",
	"#  # # #             # # !  # #    # #  ! # #  # # #        !    # #    # #    # #    # # #",
	"#  # # #           ! # #    # #    # #    # #  # # #             # #    # #    # #    # # #",
	"#  # # #             # #    # #    # #    # #  # # #             # #    # #    # #    # # #",
	"#  # # ########### # # #  # # #  # # #  # # #  # # ########### # # #  # # #  # # #  # # # #",
	"#  # # #           # # #  # # #  # # #  # # #  # # #           # # #  # # #  # # #  # # # #",
	"#                  #      #      #      #           !          #      #      #      #     #",
	"###########################################################################################",
}

type maze struct {
	width    int
	fresh    []cell
	old      []cell
	crawlers []crawler
	food     []vector
	info     string
	stopped  bool
}
type status int

func (m *maze) Nearby(count int, s vector) []vector {
	type dist struct {
		int
		vector
	}

	distances := make([]dist, len(m.food))
	for i := range distances {
		d := m.food[i].sub(s).abs()
		distances[i].int = int(math.Sqrt(float64(d.x*d.x + d.y*d.y)))
		distances[i].vector = m.food[i]
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].int < distances[j].int
	})

	if count >= len(m.food) {
		count = len(m.food)
	}

	out := make([]vector, count)

	for i := 0; i < count; i++ {
		out[i] = distances[i].vector
	}
	return out

}

func (m *maze) Path(source vector, sink vector) (path vlist) {

	visited := make(map[vector]vector)
	frontier := []vector{source}
	increment := 0

	var current vector

	for current != sink && len(frontier)-increment > 0 {
		m.Write(current, cloud)

		current = frontier[increment]
		increment++

		for _, neighbor := range m.neighbors(current) {
			if _, ok := visited[neighbor]; !ok {

				frontier = append(frontier, neighbor)
				//mark as touched
				visited[neighbor] = current

				if neighbor == sink {
					current = neighbor
					break
				}
			}
		}

	}

	if current == sink {
		for current != source {
			path.PushFront(current)
			current = visited[current]
		}
		path.PushFront(current)

		return path
	}

	return vlist{}

	//panic("no path")

}

func (m *maze) neighbors(of vector) []vector {
	neighbors := make([]vector, 4)
	check := func(at vector) bool {
		return at.x < m.width && at.x >= 0 && at.y < len(m.old)/m.width && at.y >= 0 && !m.old[at.y*m.width+at.x].Has(wall)
	}
	i := 0
	var n vector
	if n = of.add(vector{1, 0}); check(n) {
		neighbors[i] = n
		i++
	}
	if n = of.add(vector{-1, 0}); check(n) {
		neighbors[i] = n
		i++
	}
	if n = of.add(vector{0, 1}); check(n) {
		neighbors[i] = n
		i++
	}
	if n = of.add(vector{0, -1}); check(n) {
		neighbors[i] = n
		i++
	}

	return neighbors[:i]

}

func (m *maze) Convert(from []string) error {
	rand.Seed(time.Now().UnixNano())

	if len(from) < 1 {
		return errors.New("0 length input!")
	}

	m.width = len(from[0])

	for y, line := range from {
		if len(line) != m.width {
			return errors.New("uneven widths!")
		}

		m.old = append(m.old, make([]cell, len(line))...)
		m.fresh = append(m.fresh, make([]cell, len(line))...)

		for x, letter := range line {
			switch letter {
			case ' ':
			case '#':
				m.fresh[y*m.width+x].Write(wall)
			case 'X':
				m.fresh[y*m.width+x].Write(body_c)
				m.crawlers = append(m.crawlers, Crawler(m, x, y))
			case '!':
				m.fresh[y*m.width+y].Write(food)
				m.food = append(m.food, vector{x, y})
			default:
				return errors.New("unknown symbol " + string(letter))

			}

		}
	}
	//m.ScatterCrawlers(10)
	m.ScatterFood(10)

	m.info = cell(all).Enumerate()

	return nil

}
func (m *maze) ScatterFood(prob int) {
	for i := range m.fresh {
		if m.old[i].Some(wall|body_c) == 0 {
			if rand.Int()%len(m.fresh) <= prob {
				m.old[i].Write(food)
				m.food = append(m.food, vector{i % m.width, i / m.width})
			}
		}
	}
}

func (m *maze) ScatterCrawlers(prob int) {
	for i := range m.fresh {
		if m.old[i].Some(wall|body_c) == 0 {
			if rand.Int()%len(m.fresh) <= prob {
				m.fresh[i].Write(body_c)
				m.crawlers = append(m.crawlers, Crawler(m, i%m.width, i/m.width))
			}
		}
	}
}
func (m *maze) Update() {
	for i := range m.crawlers {
		m.crawlers[i].Update()
	}
	m.updateFood()
}

func (m *maze) CelltoRune(c cell, offset int) rune {
	switch {
	case c.Has(wall):
		return block
	case c.Has(food):
		return fmark
	case c.Has(body_c):
		return body_cm
	case c.Has(body_e):
		return body_em
	case c.Has(body_w):
		return body_wm
	case c.Has(body_r):
		return body_rm
	case c.Has(tendril):
		return tendrils[m.orient(tendril|body_c, offset)]
	case c.Has(cloud):
		return space
	case c == 0:
		return space
	default:
		return '?'
	}
}

//var walls = map[uint8]rune{
//	none: 'O',
//	n:    '\u2568',
//	s:    '\u2565',
//	ns:   '\u2551',
//	e:    '\u2578',
//	ne:   '\u2517',
//	se:   '\u2554',
//	nse:  '\u2560',
//	w:    '\u257A',
//	nw:   '\u255D',
//	sw:   '\u2557',
//	nsw:  '\u2563',
//	ew:   '\u2550',
//	new:  '\u2569',
//	sew:  '\u2566',
//	nsew: '\u256C',
//}
var tendrils = map[uint8]rune{
	none: ':',
	n:    '\u2575',
	s:    '\u2577',
	ns:   '\u2502',
	e:    '\u2574', //
	ne:   '\u2570', //
	se:   '\u256D', //
	nse:  '\u251C', //
	w:    '\u2576', //
	nw:   '\u256F', //
	sw:   '\u256E', //
	nsw:  '\u2524', //
	ew:   '\u2500',
	new:  '\u2534', //
	sew:  '\u252C',
	nsew: '\u253C',
}

const (
	none = 0x00
	n    = 0x01
	s    = 0x02
	ns   = 0x03
	w    = 0x04
	ne   = 0x05
	se   = 0x06
	nse  = 0x07
	e    = 0x08
	nw   = 0x09
	sw   = 0x0A
	nsw  = 0x0B
	ew   = 0x0C
	new  = 0x0D
	sew  = 0x0E
	nsew = 0x0F
)

func (m *maze) orient(o object, offset int) uint8 {
	//   x
	// x   x
	//   x
	//16 possible combinations:
	var key uint8

	dirs := []uint{
		uint(offset - m.width), //up
		uint(offset + m.width), //down
		uint(offset + 1),       //right
		uint(offset - 1),       //left
	}

	for i, dir := range dirs {
		if dir >= 0 && dir < uint(len(m.old)) {
			if m.old[dir].Some(o) != 0 {
				key |= 1 << uint(i)
			}
		}
	}
	return key

}
func (m *maze) Swap() {
	m.fresh, m.old = m.old, m.fresh
}
func (m *maze) Draw() {

	termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)

	for _, f := range m.food {
		m.fresh[f.y*m.width+f.x].Write(food)
	}
	for i := range m.crawlers {
		m.crawlers[i].Draw()
	}

	var empty cell

	for i, cel := range m.old {

		r, f, b := m.DrawCell(cel, i)

		termbox.SetCell(i%m.width, i/m.width, r, f, b)

		if cel.Has(wall) {
			m.fresh[i].Write(wall)
		}

	}
	for i := range m.old {
		m.old[i] = empty
	}

	y := 0
	x := m.width
	for _, letter := range m.info {
		if letter == '\n' {
			x = m.width
			y++
		}
		x++
		termbox.SetCell(x, y, letter, termbox.ColorDefault, termbox.ColorBlack)
	}

}
func (m *maze) DrawCell(c cell, offset int) (letter rune, fg termbox.Attribute, bg termbox.Attribute) {
	letter = m.CelltoRune(c, offset)
	fg = termbox.ColorDefault
	if c.Has(cloud) {
		bg = termbox.ColorBlue
	} else {
		bg = termbox.ColorBlack
	}
	return

}

func (m *maze) Query(x, y int) {
	if y*m.width+x < len(m.fresh) {
		m.info = m.Read(vector{x, y}).Enumerate()
	}
}
func (m *maze) updateFood() {
	if rand.Int()%5 == 0 {
		m.ScatterFood(0)
		//if rand.Int()%10 == 0 {
		//	m.ScatterCrawlers(0)
		//}
	}
	temp := make([]vector, 0, len(m.food))
	for _, f := range m.food {
		if m.old[f.y*m.width+f.x].Has(food) {
			temp = append(temp, f)
		}
	}
	m.food = temp

}
func (m *maze) Write(at vector, o object) {
	m.fresh[at.y*m.width+at.x].Write(o)
}
func (m *maze) Erase(at vector, o object) {
	m.old[at.y*m.width+at.x].Erase(o)
}
func (m *maze) Read(at vector) cell {
	return m.old[at.y*m.width+at.x]
}
func (m *maze) set(x, y int, r rune) {
	if x >= m.width || x < 0 || y >= len(m.fresh)/m.width || y < 0 {
		return
	}
	termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorBlack)
}

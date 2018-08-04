package main

import (
	"math"
	"math/rand"
)

const (
	crawlerArmNo = 5
)

type crawler struct {
	loc    vector
	arms   []arm
	prim   int
	maxLen int
	state  bstate
	enviro *maze
}

type bstate int

const (
	wandering bstate = iota
	chasing
	reaching
	eating
)

func Crawler(m *maze, x, y int) (c crawler) {
	return crawler{
		vector{x, y},
		make([]arm, 5),
		-1,
		15,
		wandering,
		m,
	}

}

func (c *crawler) Update() {
	switch c.state {
	case wandering:
		c.assign()
		c.prim = c.poll()
		for i := range c.arms {
			c.arms[i].updateLength()
		}
		if c.prim != -1 {
			c.state = reaching
		}

	case reaching:
		for i := range c.arms {
			c.arms[i].updateLength()
		}

		if c.arms[c.prim].segments.Len() > c.maxLen || c.arms[c.prim].segments.Len() == c.arms[c.prim].path.Len() {
			c.state = chasing
		}

	case chasing:
		for i := range c.arms {
			if i != c.prim {
				c.arms[i].state = retracting
			}
			c.arms[i].updateLength()
		}
		if c.arms[c.prim].segments.Len() == 0 || c.arms[c.prim].path.Len() == 0 {
			c.state = eating
			return
		}
		c.loc = c.arms[c.prim].path.PopFront()
		c.arms[c.prim].segments.PopFront()
		c.move(c.loc)

	case eating:
		c.consume()
		c.prim = -1
		c.state = wandering
	}

}

func (c *crawler) getPaths(count int) []vlist {
	nearbyFood := c.enviro.Nearby(count, c.loc)
	out := make([]vlist, len(nearbyFood))
	for i, v := range nearbyFood {
		out[i] = c.enviro.Path(c.loc, v)
	}
	return out
}

func (c *crawler) Draw() {
	for i := range c.arms {
		c.arms[i].draw(c.enviro)
	}
	c.enviro.Write(c.loc, c.debug())
}
func (c *crawler) debug() object {
	switch c.state {
	case wandering:
		return body_w
	case chasing:
		return body_c
	case reaching:
		return body_r
	case eating:
		return body_e
	}
	return 0

}

//TODO make it assign multiple paths
func (c *crawler) assign() {

	var list []int

	for i := range c.arms {
		path := c.arms[i].path
		if (path.Len() <= 0 || !c.enviro.Read(*path.I(path.Len() - 1)).Has(food)) && len(c.enviro.food) > 0 {
			list = append(list, i)

		}
	}

	paths := c.getPaths(len(list))

	for i, path := range paths {
		c.arms[list[i]].assign(path)
	}
}

func (c *crawler) move(newLoc vector) {
	for i := range c.arms {
		if i != c.prim {
			c.arms[i].move(newLoc)
		}
	}
}

func (c *crawler) consume() {
	c.enviro.Erase(c.loc, food)
}

func (c *crawler) poll() int {
	pathLength := math.MaxInt64
	prim := -1
	for i := range c.arms {
		plen := c.arms[i].path.Len()
		if pathLength > plen && rand.Int()%3 == 0 {
			prim = i
			pathLength = plen
		}
	}
	return prim
}

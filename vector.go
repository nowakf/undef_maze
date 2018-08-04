package main

import (
	"fmt"
	"math"
)

type vector struct {
	x, y int
}

func (c vector) String() string {
	return fmt.Sprintf("v(%d, %d)", c.x, c.y)
}
func (c vector) sub(o vector) vector {
	return vector{
		c.x - o.x,
		c.y - o.y,
	}
}

func (c vector) unit() vector {
	un := func(a int) int {
		if a == 0 {
			return a
		}
		return a / int(math.Abs(float64(a)))
	}
	return vector{
		un(c.x),
		un(c.y),
	}

}
func (c vector) dot(to vector) vector {
	return vector{
		c.x * to.x,
		c.y * to.y,
	}
}
func (c vector) add(to vector) vector {
	return vector{
		c.x + to.x,
		c.y + to.y,
	}
}
func (c vector) abs() vector {
	return vector{
		int(math.Abs(float64(c.x))),
		int(math.Abs(float64(c.y))),
	}
}

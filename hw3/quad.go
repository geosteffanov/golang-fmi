package main

import "github.com/fmi/go-homework/geom"

type Quad struct {
	a geom.Vector
	b geom.Vector
	c geom.Vector
	d geom.Vector
}

func NewQuad(a, b, c, d geom.Vector) Quad {
	return Quad{
		a: a,
		b: b,
		c: c,
		d: d,
	}
}

func (q *Quad) Intersect(ray geom.Ray) bool {
	alphaTriangle := Triangle{
		a: q.a,
		b: q.b,
		c: q.d,
	}

	gammaTriangle := Triangle{
		a: q.b,
		b: q.c,
		c: q.d,
	}

	return alphaTriangle.Intersect(ray) || gammaTriangle.Intersect(ray)

}

package main

import (
	"github.com/fmi/go-homework/geom"
	"math"
)

func multiplyByScalar(t float64, a geom.Vector) geom.Vector {
	return geom.Vector{
		X: t * a.X,
		Y: t * a.Y,
		Z: t * a.Z,
	}
}

func vectorDiff(a, b geom.Vector) geom.Vector {
	return geom.Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func vectorSum(a, b geom.Vector) geom.Vector {
	return geom.Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func dotProduct(a, b geom.Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func crossProduct(a, b geom.Vector) geom.Vector {
	xDet := a.Y*b.Z - a.Z*b.Y
	yDet := a.X*b.Z - a.Z*b.X
	zDet := a.X*b.Y - a.Y*b.X

	return geom.Vector{
		X: xDet,
		Y: -yDet,
		Z: zDet,
	}
}

func magnitute(a geom.Vector) float64 {
	return math.Sqrt(dotProduct(a, a))
}

func normalize(a geom.Vector) geom.Vector {
	magnitude := magnitute(a)

	return geom.Vector{
		X: a.X / magnitude,
		Y: a.Y / magnitude,
		Z: a.Z / magnitude,
	}
}

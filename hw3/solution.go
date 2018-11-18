package main

import (
	"github.com/fmi/go-homework/geom"
	"math"
)

type Triangle struct {
	a geom.Vector
	b geom.Vector
	c geom.Vector
}

func NewTriangle(a, b, c geom.Vector) Triangle {
	return Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (tr *Triangle) Intersect(ray geom.Ray) bool {
	// Ray Equation: p(t) = p0 + t.p1
	// Let n be the normal vector
	// [(p0 + p1.t) - a] * n = 0  <=> p(t) - a is parallel to the plane and hence on the plane,
	// hence p(t) is on the plane. Therefore, we need to find the value of t for which this equation
	// is satisfied.
	// The equation is equivalent to:
	//
	// t = (a * n - p0 * n) / p1 * n
	//
	//
	// But first we check if p1 * n != 0
	// If p1 * n == 0 => The ray is parallel to the plane, and hence can't intersect the triangle
	//
	// Afterwards we find the intersection point using t

	// Get the normal vector to the plane of the triangle
	normalVector := tr.normalVector()

	// Get the cosine of the angle between the normal vector and the direction of the ray
	cosAngleDirectionNormal := dotProduct(ray.Direction, normalVector)

	// Check if ray is parallel to triangle plane
	if cosAngleDirectionNormal == 0 {
		return false
	}

	// ray is not parallel to plane
	// => ray intersects plane
	// find intersection point
	// t is the value of the parameter t in the equation for the ray
	t := dotProduct(vectorDiff(tr.a, ray.Origin), normalVector) / cosAngleDirectionNormal

	// If t is less than zero then we are moving in the opposite direction of the ray
	// so we can't possibly intersect with the triangle
	if t < 0 {
		return false
	}

	intersectionPoint := vectorSum(ray.Origin, multiplyByScalar(t, ray.Direction))

	// check if intersection point is inside triangle
	return tr.inside(intersectionPoint)

}

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

type Sphere struct {
	center geom.Vector
	radius float64
}

func NewSphere(origin geom.Vector, r float64) Sphere {
	return Sphere{
		center: origin,
		radius: r,
	}
}

func (s *Sphere) Intersect(ray geom.Ray) bool {
	// Ray's vector equation
	// P=P0 + P1 . t
	// Sphere's vector equation
	// (P-C)*(P-C) - r^2 = 0

	// Substituting the P from the first one into the second one we get

	// P1^2 * t^2 + 2P1*(P0-C)*t +(P0-C)^2 - r^2 = 0

	// A quadratic equation of t of the form at^2 + 2bt + c = 0
	// Our solution depends only on the sign of the determinant
	// And the signs of the roots if they are real
	a := dotProduct(ray.Direction, ray.Direction)

	displacement := vectorDiff(ray.Origin, s.center)
	b := 2 * dotProduct(ray.Direction, displacement)

	c := dotProduct(displacement, displacement) - s.radius*s.radius

	determinant := b*b - 4*a*c

	if determinant < 0 {
		// Determinant is less than zero. Therefore, there are no real roots
		// Hence, no t for which the ray intersects the sphere
		return false
	}

	// Determinant is nonnegative, therefore there is a real solution
	// Since the ray has an Origin and we consider only t >= 0 as a value
	// We need to check if there is a positive root

	// We will use Vieta's formulas
	// t1 * t2 = c/a
	// t + t2 = -b/a
	rootProduct := c / a

	if rootProduct <= 0 {
		// rootProduct is less than or equal to zero
		// this means that the origin is either an intersection point
		// or there is a positive t which is on the sphere
		return true
	}

	// rootProduct is > 0 => Both roots have the same sign. We just need to check whether
	// it is positive
	positive := -b/a > 0

	return positive
}

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

func determinant(mat [][]float64) float64 {
	return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
}

func (tr *Triangle) normalVector() geom.Vector {
	left := vectorDiff(tr.c, tr.a)
	right := vectorDiff(tr.b, tr.a)

	normalVector := normalize(crossProduct(left, right))

	return normalVector
}

func (tr *Triangle) inside(vector geom.Vector) bool {
	// Find whether a point is inside a triangle using its barycentric coordinates.
	pMinusA := vectorDiff(vector, tr.a)
	bMinusA := vectorDiff(tr.b, tr.a)
	cMinusA := vectorDiff(tr.c, tr.a)

	// Here we will use Cramer's Rule to find out the solutions
	// to the system of equations
	matrix := [][]float64{
		{bMinusA.X, cMinusA.X},
		{bMinusA.Y, cMinusA.Y},
	}

	a1 := [][]float64{
		{pMinusA.X, cMinusA.X},
		{pMinusA.Y, cMinusA.Y},
	}

	a2 := [][]float64{
		{bMinusA.X, pMinusA.X},
		{bMinusA.Y, pMinusA.Y},
	}

	beta := determinant(a1) / determinant(matrix)
	gamma := determinant(a2) / determinant(matrix)

	if beta < 0 || gamma < 0 {
		return false
	}

	if beta > 1 || gamma > 1 {
		return false
	}

	if beta+gamma > 1 {
		return false
	}

	return true
}

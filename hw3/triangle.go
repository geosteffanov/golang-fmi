package main

import (
	"github.com/fmi/go-homework/geom"
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
	// Ray Equation: P = P0 + P1 . t
	// Let n be the normal vector
	// [(P0 + P1.t) - A] * n = 0  <=> P(t) - A is parallel to the plane and hence on the plane,
	// hence P(t) is on the plane. Therefore, we need to find the value of t for which this equation
	// is satisfied.
	// The equation is equivalent to:
	//
	// t = (A * n - P0 * n) / P1 * n
	//
	//
	// But first we check if P1 * n != 0
	// If P1 * n == 0 => The ray is parallel to the plane, and hence can't intersect the triangle
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
	t := dotProduct(vectorDiff(ray.Origin, tr.a), normalVector)

	intersectionPoint := vectorSum(ray.Origin, multiplyByScalar(t, ray.Direction))

	// check if intersection point is inside triangle
	return tr.inside(intersectionPoint)

}

func (tr *Triangle) normalVector() geom.Vector {
	left := vectorDiff(tr.c, tr.a)
	right := vectorDiff(tr.b, tr.a)

	normalVector := normalize(crossProduct(left, right))

	return normalVector
}

// More documentation on this will follow.
func (tr *Triangle) inside(vector geom.Vector) bool {
	// Find whether a point is inside a triangle using its barycentric coordinates.
	pMinusA := vectorDiff(vector, tr.a)
	bMinusA := vectorDiff(tr.b, tr.a)
	cMinusA := vectorDiff(tr.c, tr.a)

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

func determinant(mat [][]float64) float64 {
	return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
}

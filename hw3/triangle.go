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

	if !tr.rayPointsTowardsPlane(ray) {
		return false
	}

	// ray is not parallel to plane
	// => ray intersects plane
	// find intersection point
	// t is the value of the parameter t in the equation for the ray
	t := dotProduct(vectorDiff(tr.a, ray.Origin), normalVector)


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

// This function returns the distance from a point to the plane
// in which the triangle lies.
func (tr *Triangle) distanceFromPointToPlane(point geom.Vector) float64 {
	// Given a plane equation: Ax + By + Cz + D = 0
	// and a point <x1,y1,z1> in space, the distance
	// from the point to the plane can be found by
	// the following formula:
	//
	// d = | Ax1 + By1 + Cz1 + D | /  sqrt(A^2 + B^2+C^2)
	//

	// First we calculate A,B,C,D

	// The plane is defined as (P - A) * n = 0
	// where n is the normal unit vector of the plane
	// and A is one of the vertices of the triangle.

	// If we write down P = <x,y,z>
	// and A = <xA,yA,zA>, and n = <N1, N2, N3>
	// P * n - A*n = 0 is therefore equivalent to
	//
	// N1.x + N2.y + N3.z - A*n = 0
	// Hence A=N1, B=N2, C=N3, D = -A*n
	// from which we find A,B,C,D respectively.

	// Since d = | Ax1 + By1 + Cz1 + D | /  sqrt(A^2 + B^2+C^2)
	// we see that we can actually write this down as

	// d = | n * <x1, y1, z1>  - A*n | / |n|,
	// and since we use a normalized n
	// |n| = 1
	// hence,
	// d = | n * (<x1, y1,z1> - A) |

	normalVector := tr.normalVector()
	//D := -1 * dotProduct(normalVector, tr.a)

	diff := vectorDiff(point, tr.a)

	return math.Abs(dotProduct(normalVector, diff))

}

// utility function that checks whether a given ray points outwards of a plane
// if a ray points outwards, then it can't possibly intersect the plane
func (tr *Triangle) rayPointsTowardsPlane(ray geom.Ray) bool {
	// if the origin is on the plane we return true
	if tr.inPlane(ray.Origin) {
		return true
	}

	distanceFromOriginToPlane := tr.distanceFromPointToPlane(ray.Origin)
	normalizedDirection := normalize(ray.Direction)
	halfDistance := distanceFromOriginToPlane / 2
	reducedDirection := multiplyByScalar(halfDistance, normalizedDirection)

	smallIncrement := vectorSum(ray.Origin, reducedDirection)
	distanceFromIncrementedPoint := tr.distanceFromPointToPlane(smallIncrement)

	return distanceFromIncrementedPoint < distanceFromOriginToPlane
}

// point is on the plane of the triangle
func (tr *Triangle) inPlane(point geom.Vector) bool {
	return dotProduct(vectorDiff(point, tr.a), tr.normalVector()) == 0
}

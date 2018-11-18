package main

import "github.com/fmi/go-homework/geom"

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

package main

import (
	"fmt"
	"github.com/fmi/go-homework/geom"
	"github.com/gsamokovarov/assert"
	"testing"
)

func TestTriangle_Intersect(t *testing.T) {
	tests := []struct {
		triangle   Triangle
		ray        geom.Ray
		intersects bool
	}{
		{
			triangle: Triangle{
				a: geom.Vector{0, 0, 1},
				b: geom.Vector{0, 1, 1},
				c: geom.Vector{1, 0, 1},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{0, 0, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: true,
		},

		{
			triangle: Triangle{
				a: geom.Vector{2, 0, 0},
				b: geom.Vector{1, 0, 0},
				c: geom.Vector{1, 0, 1},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{0, 0, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},

		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{0, 0, 1},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{0, 0, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: true,
		},

		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{1.0 / 2, 1.0 / 2, 1},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{0, 0, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},
		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{0, 0, 0},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{2, 0, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},
		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{0, 0, 0},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{0, 2, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},
		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{0, 0, 0},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{-1, -1, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},
		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{0, 0, 0},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{1, 1, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},
		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{0, 0, 0},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{1.0 / 4, 1.0 / 4, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: true,
		},
		{
			triangle: Triangle{
				a: geom.Vector{1, 0, 0},
				b: geom.Vector{0, 1, 0},
				c: geom.Vector{0, 0, 0},
			},
			ray: geom.Ray{
				Origin:    geom.Vector{1.0 / 4, 1.0 / 4, 1.0 / 2},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := test.triangle.Intersect(test.ray)
			assert.Equal(t, test.intersects, actual)
		})
	}
}

func TestTriangle_distanceFromPoint(t *testing.T) {
	triangle := NewTriangle(geom.NewVector(1, 0, 0), geom.NewVector(0, 1, 0), geom.NewVector(0, 0, 0))

	tests := []struct {
		point   geom.Vector
		distace float64
	}{
		{
			point:   geom.Vector{X: 1, Y: 0, Z: 0},
			distace: 0,
		},
		{
			point:   geom.Vector{X: 0, Y: 0, Z: 0},
			distace: 0,
		},
		{
			point:   geom.Vector{X: 0, Y: 1, Z: 0},
			distace: 0,
		},

		{
			point:   geom.Vector{X: 100, Y: 0, Z: 0},
			distace: 0,
		},
		{
			point:   geom.Vector{X: 0, Y: 0, Z: 1},
			distace: 1,
		},
		{
			point:   geom.Vector{X: 0, Y: 0, Z: 25},
			distace: 25,
		},
		{
			point:   geom.Vector{X: 18273981, Y: 23, Z: -32},
			distace: 32,
		},
		{
			point:   geom.Vector{X: 23.2, Y: 2333.3, Z: 15.3},
			distace: 15.3,
		},
	}

	for index, test := range tests {
		desc := fmt.Sprintf("Test no: %d", index)
		t.Run(desc, func(t *testing.T) {
			actual := triangle.distanceFromPointToPlane(test.point)
			assert.Equal(t, test.distace, actual)
		})
	}
}

func TestTriangle_rayPointsTowardsPlane(t *testing.T) {
	t.Run("Specific test", func(t *testing.T) {
		triangle := Triangle{
			a: geom.Vector{0, 0, 1},
			b: geom.Vector{0, 1, 1},
			c: geom.Vector{1, 0, 1},
		}
		ray := geom.Ray{
			Origin:    geom.Vector{0, 0, 0},
			Direction: geom.Vector{0, 0, 1},
		}

		assert.True(t, triangle.Intersect(ray))

	})

	//triangle := Triangle{
	//	a: geom.Vector{X: 0, Y: 0, Z: 0},
	//	b: geom.Vector{X: 0, Y: 1, Z: 0},
	//	c: geom.Vector{X: 1, Y: 0, Z: 0},
	//}
	//
	//tests := []struct {
	//	ray                geom.Ray
	//	pointsTowardsPlane bool
	//}{
	//	{
	//		ray: geom.Ray{
	//			Origin:    geom.Vector{0, 0, 0},
	//			Direction: geom.Vector{0, 0, 1},
	//		},
	//		pointsTowardsPlane: true,
	//	},
	//	{
	//		ray: geom.Ray{
	//			Origin:    geom.Vector{0, 0, 1},
	//			Direction: geom.Vector{0, 0, 1},
	//		},
	//		pointsTowardsPlane: false,
	//	},
	//	{
	//		ray: geom.Ray{
	//			Origin:    geom.Vector{0, 0, 1},
	//			Direction: geom.Vector{0, 0, -1},
	//		},
	//		pointsTowardsPlane: true,
	//	},
	//	{
	//		ray: geom.Ray{
	//			Origin:    geom.Vector{0, 0, 1},
	//			Direction: geom.Vector{0, 1, 0},
	//		},
	//		pointsTowardsPlane: false,
	//	},
	//}
	//
	//for index, test := range tests {
	//	desc := fmt.Sprintf("test no: %d", index)
	//	t.Run(desc, func(t *testing.T) {
	//		actual := triangle.rayPointsTowardsPlane(test.ray)
	//		assert.Equal(t, test.pointsTowardsPlane, actual)
	//	})
	//}
}

package main

import (
	"fmt"
	"github.com/fmi/go-homework/geom"
	"github.com/gsamokovarov/assert"
	"testing"
)

func TestQuad_Intersect(t *testing.T) {
	quad := Quad{
		a: geom.Vector{0, 0, 0},
		b: geom.Vector{1, 0, 0},
		c: geom.Vector{1, 1, 0},
		d: geom.Vector{0, 1, 0},
	}

	tests := []struct {
		ray        geom.Ray
		intersects bool
	}{
		{
			ray: geom.Ray{
				Origin:    geom.Vector{1, 1, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{2, 2, 0},
				Direction: geom.Vector{0, 0, 1},
			},
			intersects: false,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{2, 0, 0},
				Direction: geom.Vector{0, -1, 0},
			},
			intersects: false,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{0, 0, 0.5},
				Direction: geom.Vector{0.5, 0.5, -0.5},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{1, 0.5, 0},
				Direction: geom.Vector{-1, 0, 0},
			},
			intersects: false,
		},
	}

	for index, test := range tests {
		desc := fmt.Sprintf("Test no: %d", index)
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.intersects, quad.Intersect(test.ray))
		})
	}
}

func TestSphere_Intersect(t *testing.T) {
	sphere := NewSphere(
		geom.NewVector(0, 0, 0),
		1,
	)

	tests := []struct {
		ray        geom.Ray
		intersects bool
	}{
		{
			ray: geom.Ray{
				Origin:    geom.Vector{X: 0, Y: 0, Z: 0},
				Direction: geom.Vector{X: 0, Y: 0, Z: 1},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{X: 0, Y: 1, Z: 0},
				Direction: geom.Vector{X: 0, Y: 0, Z: 1},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{X: 0, Y: 2, Z: 0},
				Direction: geom.Vector{X: 0, Y: 0, Z: 1},
			},
			intersects: false,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{X: 2, Y: 0, Z: 0},
				Direction: geom.Vector{X: -1, Y: 0, Z: 0},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{X: 2, Y: 0, Z: 0},
				Direction: geom.Vector{X: -0.5, Y: 0, Z: 0},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{X: 2, Y: 0, Z: 0},
				Direction: geom.Vector{X: 0.5, Y: 0, Z: 0},
			},
			intersects: false,
		},
		{
			ray: geom.Ray{
				Origin:    geom.Vector{X: 2, Y: 0, Z: 0},
				Direction: geom.Vector{X: 0.5, Y: 0, Z: 0},
			},
			intersects: false,
		},
	}

	for index, test := range tests {
		desc := fmt.Sprintf("Test no: %d", index)
		t.Run(desc, func(t *testing.T) {
			assert.Equal(t, test.intersects, sphere.Intersect(test.ray))
		})
	}
}

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

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
				c: geom.Vector{1.0/2, 1.0/2, 1},
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
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := test.triangle.Intersect(test.ray)
			assert.Equal(t, test.intersects, actual)
		})
	}
}

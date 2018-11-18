package main

import (
	"testing"
	"github.com/fmi/go-homework/geom"
	"fmt"
	"github.com/gsamokovarov/assert"
)

func TestQuadIntersect(t *testing.T) {
	quad := Quad{
		a: geom.Vector{0, 0, 0},
		b: geom.Vector{1, 0, 0},
		c: geom.Vector{1, 1, 0},
		d: geom.Vector{0, 1, 0},
	}

	tests := []struct{
		ray geom.Ray
		intersects bool
	}{
		{
			ray: geom.Ray{
				Origin: geom.Vector{1,1,0},
				Direction: geom.Vector{0,0,1},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin: geom.Vector{2,2,0},
				Direction: geom.Vector{0,0,1},
			},
			intersects: false,
		},
		{
			ray: geom.Ray{
				Origin: geom.Vector{2,0,0},
				Direction: geom.Vector{0,-1,0},
			},
			intersects: false,
		},
		{
			ray: geom.Ray{
				Origin: geom.Vector{0,0,0.5},
				Direction: geom.Vector{0.5,0.5,-0.5},
			},
			intersects: true,
		},
		{
			ray: geom.Ray{
				Origin: geom.Vector{1,0.5,0},
				Direction: geom.Vector{-1,0,0},
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


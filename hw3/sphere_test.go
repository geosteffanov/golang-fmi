package main

import (
	"fmt"
	"github.com/fmi/go-homework/geom"
	"github.com/gsamokovarov/assert"
	"testing"
)

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

package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	assert.Equal(t, want, got)
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		assert.Equal(t, want, got)
	}

	t.Run("rectanbles", func(t *testing.T) {
		r := Rectangle{12.0, 6.0}
		want := 12.0 * 6.0
		checkArea(t, r, want)
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		want := 314.1592653589793
		checkArea(t, c, want)
	})
}

func TestAreaTable(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectabnel", Rectangle{Width: 12, Height: 6}, 72.0},
		{"circle", Circle{Radius: 10}, 314.1592653589793},
		{"triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			assert.Equal(t, tt.want, got)
		})
	}
}

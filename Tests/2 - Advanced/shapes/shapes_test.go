package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12, "Rectangle A"}
		expectedArea := float64(120)
		outputArea := rect.Area()

		if expectedArea != outputArea {
			t.Fatalf("Expected %f received %f", expectedArea, outputArea)
		}

	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10, "Circle A"}
		expectedArea := float64(math.Pi * 100)
		outputArea := circ.Area()

		if expectedArea != outputArea {
			t.Fatalf("Expected %f received %f", expectedArea, outputArea)
		}
	})
}

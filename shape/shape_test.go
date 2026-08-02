package shape

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	// t.Run("rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	// got := rectangle.Area()
	// 	// want := 72.0

	// 	checkArea(t, rectangle, 72)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	// got := circle.Area()
	// 	// want := 314.1592653589793

	// 	checkArea(t, circle, 314.1592653589793)

	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{"circle", Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			// got := tt.shape.Area()
			// if got != tt.want {
			// 	t.Errorf("got %g want %g", got, tt.want)
			// }
			checkArea(t, tt.shape, tt.hasArea)
		})
	}

	// การบ้าน สร้าง Triangle
	// 1. สร้าง struct
	// 2. สร้าง method Area()
	// 3. สร้าง test
}

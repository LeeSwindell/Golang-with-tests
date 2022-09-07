package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{shape: Rectangle{Width: 2, Height: 4}, want: 8},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 2, Height: 5}, want: 6},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g, want %g",tt.shape, got, tt.want)
		}
	}

}
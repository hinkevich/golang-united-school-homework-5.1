package square

import "testing"

type testDataEndPoint struct {
	sq   Square
	want Point
}

func TestSquare(t *testing.T) {
	tests := []testDataEndPoint{
		testDataEndPoint{sq: Square{Point{x: 0, y: 0}, 3}, want: Point{3, -3}},
		testDataEndPoint{sq: Square{Point{x: -5, y: -6}, 3}, want: Point{-2, -9}},
		testDataEndPoint{sq: Square{Point{x: 5, y: 5}, 5}, want: Point{10, 0}},
		testDataEndPoint{sq: Square{Point{x: 10, y: -10}, 3}, want: Point{13, -13}},
		testDataEndPoint{sq: Square{Point{x: -10, y: 10}, 3}, want: Point{-7, 7}},
	}
	for i, test := range tests {
		got := test.sq.End()
		if got != test.want {
			t.Errorf("test number %d is failed, got: %d ; want: %d ", i+1, got, test.want)
		}
	}
}

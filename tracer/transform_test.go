package tracer

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	data := []struct {
		transform Mat4
		pt        Tuple
		expect    Tuple
	}{
		{I4.Translate(5, -3, 2), Point(-3, 4, 5), Point(2, 1, 7)},
		{I4.Translate(5, -3, 2).Inverse(), Point(-3, 4, 5), Point(-8, 7, 3)},
		{I4.Translate(5, -3, 2), Vector(-3, 4, 5), Vector(-3, 4, 5)},
	}
	for _, row := range data {
		result := row.transform.TimesTuple(row.pt)
		if result != row.expect {
			t.Errorf("Point %v should have translated to %v but was %v", row.pt, row.expect, result)
		}
	}
}

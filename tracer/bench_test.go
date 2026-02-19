package tracer

// benchResults holds all benchmark results to prevent compiler optimizations
type benchResults struct {
	// Matrix results
	matrix struct {
		mat2 Mat[Size2]
		mat3 Mat[Size3]
		mat4 Mat[Size4]
	}
	// Canvas results
	canvas struct {
		c   Canvas
		str string
	}
	// Color results
	color struct {
		c     LinearColor
		float float64
	}
	// Ray results
	ray struct {
		r    Ray
		ints []Intersect
	}
	// Transform results
	transform struct {
		mat4 Mat[Size4]
	}
	// Tuple results
	tuple struct {
		t     Tuple
		float float64
	}
	// Shared scalar results
	scalar struct {
		float float64
		bool  bool
	}
}

var results benchResults

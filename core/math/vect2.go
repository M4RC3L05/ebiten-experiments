package math

type Vector2 struct {
	X, Y float64
}

var (
	VECTOR_ZERO  = MakeVector2(0, 0)
	VECTOR_ONE   = MakeVector2(1, 1)
	VECTOR_UP    = MakeVector2(0, -1)
	VECTOR_DOWN  = MakeVector2(0, 1)
	VECTOR_LEFT  = MakeVector2(-1, 0)
	VECTOR_RIGHT = MakeVector2(1, 0)
)

func MakeVector2(x float64, y float64) Vector2 {
	return Vector2{x, y}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return MakeVector2(v.X+other.X, v.Y+other.Y)
}

func (v Vector2) AddScalar(n float64) Vector2 {
	return MakeVector2(v.X+n, v.Y+n)
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return MakeVector2(v.X-other.X, v.Y-other.Y)
}

func (v Vector2) SubScalar(n float64) Vector2 {
	return MakeVector2(v.X-n, v.Y-n)
}

func (v Vector2) Mult(other Vector2) Vector2 {
	return MakeVector2(v.X*other.X, v.Y*other.Y)
}

func (v Vector2) MultScalar(n float64) Vector2 {
	return MakeVector2(v.X*n, v.Y*n)
}

func (v Vector2) Equals(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

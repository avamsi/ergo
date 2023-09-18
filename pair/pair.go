package pair

type Pair[T1 any, T2 any] struct {
	first  T1
	second T2
}

func New[T1 any, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{first, second}
}

func (p Pair[T1, T2]) Unpack() (T1, T2) {
	return p.first, p.second
}

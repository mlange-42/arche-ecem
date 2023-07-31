package main

// Grid data
type Grid[T any] struct {
	Data   []T
	width  int
	height int
}

// NewGrid creates a new grid
func NewGrid[T any](width, height int) Grid[T] {
	gr := make([]T, width*height)
	return Grid[T]{
		Data:   gr,
		width:  width,
		height: height,
	}
}

func (g *Grid[T]) Width() int {
	return g.width
}

func (g *Grid[T]) Height() int {
	return g.height
}

func (g *Grid[T]) Contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < g.width && y < g.height
}

func (g *Grid[T]) Get(x, y int) T {
	return g.Data[x+y*g.width]
}

func (g *Grid[T]) GetPointer(x, y int) *T {
	return &g.Data[x+y*g.width]
}

func (g *Grid[T]) Set(x, y int, value T) {
	g.Data[x+y*g.width] = value
}

func (g *Grid[T]) Fill(value T) {
	for i := range g.Data {
		g.Data[i] = value
	}
}

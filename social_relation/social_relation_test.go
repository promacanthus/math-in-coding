package socialrelation

import "testing"

func TestBFS(t *testing.T) {
	BFS(generate(100, 1000), 7)
}

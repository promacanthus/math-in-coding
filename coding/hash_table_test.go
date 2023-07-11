package coding

import "testing"

func Test_container(t *testing.T) {
	c := NewContainer()
	c.Insert(1)
	c.GetRandom()
	c.Remove(1)
}

package name

import (
	"bytes"
)

// name Component is a string
type Component string

//converts a slice of bytes to a string and returns a component with that value
func ComponentFromBytes(b []byte) Component {
	return Component(string(b))
}

//converts a component to a slice of bytes
func (c Component) ComponentToBytes() []byte {
	return []byte(c)
}

//returns a Component with the given value in s
func ComponentFromString(s string) Component {
	return Component(s)
}

//returns the value stored in the component c
func (c Component) ToString() string {
	return string(c)
}

// copies the caller's value into a new component and returns it
func (c Component) Copy() Component {
	return c
}

//Compares 2 components and returns (0: if equal, -1: if c < other, and +1: if c > other)
func (c Component) Compare(other Component) int {
	return bytes.Compare([]byte(c), []byte(other))
}

// returns a boolean instead of an integer
func (c Component) Equals(other Component) bool {
	return c == other
}

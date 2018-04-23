package name

import (
	"fmt"
	"strings"
)

type Name []Component

//creates a new name from a number of components
func NewName(cs ...Component) Name {
	return Name(cs)
}

func NewNameFromString(s string) Name {
	components := strings.Split("/", s)
	cmp := make([]Component, len(components), len(components))
	for i, _ := range components {
		cmp[i] = Component(components[i])
	}
	fmt.Printf("%v", NewName(cmp...))
	return NewName(cmp...)
}

//returns the number of components in a name
func (n Name) Size() int {
	return len(n)
}

//checks if the name is empty
func (n Name) IsEmpty() bool {
	return n.Size() == 0
}

//in case of an error and we got a negative index turn it into a positive one, witout throuing an error
func (n Name) positiveIndex(i int) int {
	if i < 0 {
		i = -1 * i
	}
	return i
}

//returns the subname of --count-- components starting at offset
func (n Name) GetSubName(offset, count int) Name {
	start := n.positiveIndex(offset)
	end := start + count
	if end > n.Size() {
		end = n.Size()
	}
	return n[start:end]
}

//returns the prefix of the first --count-- Component
func (n Name) GetPrefix(count int) Name {
	return n.GetSubName(0, n.positiveIndex(count))
}

//returns the last --count-- components
func (n Name) GetSuffix(count int) Name {
	count = n.positiveIndex(count)
	start := n.Size() - count
	if start < 0 {
		start = 0
	}
	return n.GetSubName(start, count)
}

//compares tow names, returns ( 0: if equal, -1: if n < other, 1: if n > other )
func (n Name) Compare(other Name) int {
	nSize := n.Size()
	otherSize := other.Size()

	if nSize < otherSize {
		return -1
	} else if otherSize < nSize {
		return 1
	}
	//keep going while the components are the same, return the result if they are different
	for i := 0; i < nSize && i < otherSize; i++ {
		result := n[i].Compare(other[i]) //comparing components --not recursive--
		if result != 0 {
			return result
		}
	}
	return 0
}

//returns boolean instead of int
func (n Name) Equals(other Name) bool {
	return n.Compare(other) == 0
}

//returns the name in a from of a string
func (n Name) ToString() string {
	stringComponents := []string{}
	for _, c := range n {
		stringComponents = append(stringComponents, c.ToString())
	}
	return fmt.Sprintf("/%s", strings.Join(stringComponents, "/"))
}

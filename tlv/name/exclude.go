package name

type Exclude []Component

//the type any is used to match anything
var Any = Component("")

//creates a new name from a number of components
func NewExclude(cs ...Component) Exclude {
	return Exclude(cs)
}

func (e Exclude) Matches(c Component) bool {
	for i := 0; i < len(e); i++ {
		if e[i] == Any {
			//skip this component
		} else {
			if c.Equals(e[i]) {
				return true
			}
		}
	}
	return false
}

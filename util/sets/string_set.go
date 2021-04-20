package sets

type void struct{}

var member void

type StringSet map[string]void

func NewStringSet() StringSet {
	return make(map[string]void)
}

func (set *StringSet) Add(item string) {
	(*set)[item] = member
}

func (set *StringSet) Remove(item string) {
	delete(*set, item)
}

func (set *StringSet) Exists(item string) bool {
	_, exists := (*set)[item]
	return exists
}

func (set *StringSet) ToSlice() []string {
	var slice []string
	for item := range *set {
		slice = append(slice, item)
	}
	return slice
}

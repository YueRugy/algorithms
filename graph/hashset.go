package graph

const defaultCap = 1 << 4

type HashSet struct {
	array map[string]struct{}
}

func (hs *HashSet) Add(str string) bool {
	if _, ok := hs.array[str]; ok != true {
		return ok
	}
	hs.array[str] = struct{}{}
	return true
}

func (hs *HashSet) Contains(str string) bool {
	_, ok := hs.array[str]
	return ok
}

func (hs *HashSet) Clear() {
	hs.array = make(map[string]struct{})
}

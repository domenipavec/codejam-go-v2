package stringmap

type StringMap struct {
	Mapping map[string]int
	Strings []string
}

func New() *StringMap {
	return &StringMap{
		Mapping: make(map[string]int),
		Strings: make([]string, 0, 1000),
	}
}

func (sm *StringMap) Int(str string) int {
	index, ok := sm.Mapping[str]
	if ok {
		return index
	}

	index = len(sm.Strings)
	sm.Mapping[str] = index
	sm.Strings = append(sm.Strings, str)
	return index
}

func (sm *StringMap) Get(i int) string {
	return sm.Strings[i]
}

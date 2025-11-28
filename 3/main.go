package main

type StringIntMap struct {
	m map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		m: make(map[string]int),
	}
}

func (s *StringIntMap) Add(key string, value int) {
	// TODO
}

func (s *StringIntMap) Remove(key string) {
	// TODO
}

func (s *StringIntMap) Copy() map[string]int {
	// TODO
}

func (s *StringIntMap) Exists(key string) bool {
	// TODO
}

func (s *StringIntMap) Get(key string) (int, bool) {
	// TODO
}

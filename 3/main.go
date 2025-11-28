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
	s.m[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.m, key)
}

func (s *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for k, v := range s.m {
		copyMap[k] = v
	}
	return copyMap
}

func (s *StringIntMap) Exists(key string) bool {
	_, exists := s.m[key]
	return exists
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, exists := s.m[key]
	return value, exists
}

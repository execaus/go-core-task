package main

type StringIntMap struct {
	m map[string]int
}

// NewStringIntMap создает и возвращает новый экземпляр StringIntMap.
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		m: make(map[string]int),
	}
}

// Add добавляет ключ и значение в карту.
func (s *StringIntMap) Add(key string, value int) {
	s.m[key] = value
}

// Remove удаляет ключ и соответствующее значение из карты.
func (s *StringIntMap) Remove(key string) {
	delete(s.m, key)
}

// Copy создает и возвращает копию внутренней карты.
func (s *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for k, v := range s.m {
		copyMap[k] = v
	}
	return copyMap
}

// Exists проверяет, существует ли ключ в карте.
func (s *StringIntMap) Exists(key string) bool {
	_, exists := s.m[key]
	return exists
}

// Get возвращает значение по ключу и булево значение, указывающее, существует ли ключ.
func (s *StringIntMap) Get(key string) (int, bool) {
	value, exists := s.m[key]
	return value, exists
}

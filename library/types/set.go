package types

import "sync"

type Void struct {
}
type StringSet map[string]Void

func (s *StringSet) put (key string) {
	(*s)[key] = Void{}
}
func (s *StringSet) remove (key string) {
	delete(*s, key)
}
func (s *StringSet) contains (key string) bool {
	_, ok := (*s)[key]
	return ok
}

type SafeStringSet struct {
	StringSet
	sync.Mutex
}

func (s *SafeStringSet) Put (key string) {
	s.Lock()
	defer s.Unlock()

	s.StringSet.put(key)

}
func (s *SafeStringSet) Remove (key string) {
	s.Lock()
	defer s.Unlock()
	s.StringSet.remove(key)
}
func (s *SafeStringSet) Contains (key string) bool {
	s.Lock()
	defer s.Unlock()
	return s.StringSet.contains(key)
}
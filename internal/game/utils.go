package game

import "fmt"

type BodySet struct {
	elems map[string]struct{}
}

func NewSet(coords ...Coord) *BodySet {
	set := BodySet{elems: map[string]struct{}{}}

	return set.add(coords...)
}

func (s *BodySet) add(coords ...Coord) *BodySet {
	for _, c := range coords {
		key := fmt.Sprintf("%d:%d", c.X, c.Y)
		s.elems[key] = struct{}{}
	}

	return s
}

func (s *BodySet) exits(coord Coord) bool {
	key := fmt.Sprintf("%d:%d", coord.X, coord.Y)
	_, ok := s.elems[key]

	return ok
}

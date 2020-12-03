package slope

import "fmt"

type SlopeWalker struct {
	slope    *SlopeMap
	row, col int
}

func NewWalker(s *SlopeMap) *SlopeWalker {
	return &SlopeWalker{
		slope: s,
	}
}

func (s *SlopeWalker) Walk() error {
	return s.WalkWith(3, 1)
}

func (s *SlopeWalker) WalkWith(right, down int) error {
	if len(s.slope)-1 >= s.row+down {
		s.row += down
	} else {
		return fmt.Errorf("Cannot move to row %d", s.row+down)
	}

	s.col = (s.col + right) % len(s.slope[s.row])

	return nil
}

func (s *SlopeWalker) Current() MapMarker {
	return s.slope.MarkerAt(s.row, s.col)
}

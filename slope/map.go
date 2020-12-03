package slope

import "fmt"

type SlopeMap [323][31]MapMarker

func (s *SlopeMap) PrintMap() {
	for _, x := range s {
		for _, y := range x {
			fmt.Printf("%s", y)
		}
		fmt.Printf("\n")
	}
}

func (s *SlopeMap) MarkerAt(x, y int) MapMarker {
	return s[x][y]
}

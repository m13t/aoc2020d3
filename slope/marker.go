package slope

type MapMarker rune

const (
	MapTree MapMarker = '#'
	MapFree MapMarker = '.'
)

func (mm MapMarker) String() string {
	switch mm {
	case MapTree:
		return "#"
	case MapFree:
		return "."
	}

	return string(rune(mm))
}

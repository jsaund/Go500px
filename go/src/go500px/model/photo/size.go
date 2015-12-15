package photo

// Size defines the dimensions of the image in pixels.
type Size interface {
	Width() int
	Height() int
}

// SizeImpl implements the Size interface.
type SizeImpl struct {
	width  int
	height int
}

var idToSize = map[int]Size{
	1:   NewCroppedSize(70),
	2:   NewCroppedSize(140),
	3:   NewCroppedSize(280),
	100: NewCroppedSize(100),
	200: NewCroppedSize(200),
	440: NewCroppedSize(440),
	600: NewCroppedSize(600),
}

// NewSizeFromID constructs a new Size object from a 500px API image size ID.
func NewSizeFromID(id int) Size {
	s, ok := idToSize[id]
	if ok {
		return s
	}
	return idToSize[1]
}

// NewCroppedSize constructs a new Size object which represents a cropped image.
// The image will have the same width and height dimensions.
func NewCroppedSize(size int) *SizeImpl {
	return &SizeImpl{size, size}
}

// Width implements the Size interface. Returns the image width in pixels.
func (s *SizeImpl) Width() int {
	return s.width
}

// Height implements the Size interface. Returns the image height in pixels.
func (s *SizeImpl) Height() int {
	return s.height
}

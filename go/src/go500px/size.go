package go500px

// Size defines the dimensions of the image in pixels.
type Size interface {
	Width() int32
	Height() int32
}

// SizeImpl implements the Size interface.
type size struct {
	width  int32
	height int32
}

var idToSize = map[int32]Size{
	int32(1):   NewCroppedSize(70),
	int32(2):   NewCroppedSize(140),
	int32(3):   NewCroppedSize(280),
	int32(100): NewCroppedSize(100),
	int32(200): NewCroppedSize(200),
	int32(440): NewCroppedSize(440),
	int32(600): NewCroppedSize(600),
}

// NewSizeFromID constructs a new Size object from a 500px API image size ID.
func NewSizeFromID(id int32) Size {
	s, ok := idToSize[id]
	if ok {
		return s
	}
	return idToSize[1]
}

// NewCroppedSize constructs a new Size object which represents a cropped image.
// The image will have the same width and height dimensions.
func NewCroppedSize(s int32) *size {
	return &size{s, s}
}

// Width implements the Size interface. Returns the image width in pixels.
func (s *size) Width() int32 {
	return s.width
}

// Height implements the Size interface. Returns the image height in pixels.
func (s *size) Height() int32 {
	return s.height
}

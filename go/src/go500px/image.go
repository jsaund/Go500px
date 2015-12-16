package go500px

// Image provides the client with the necessary information to retrieve and display an image.
type Image interface {
	URL() string
	Size() Size
	Format() string
	Equals(other Image) bool
}

// ImageImpl implements the Image interface.
type image struct {
	ImageURL    string `json:"https_url"`
	SizeID      int32  `json:"size"`
	ImageFormat string `json:"format"`
	size        Size
}

type Images struct {
	images []*image
}

func NewImages() *Images {
	return new(Images)
}

func (i *Images) Count() int {
	return len(i.images)
}

func (i *Images) Get(index int) Image {
	return i.images[index]
}

// URL implements the Image interface. Returns the image URL to download the image.
func (i *image) URL() string {
	return i.ImageURL
}

// Size implements the Image interface. Returns the Size of the image in pixels.
func (i *image) Size() Size {
	if i.size == nil {
		i.size = NewSizeFromID(i.SizeID)
	}
	return i.size
}

// Format implements the Image interface. Returns the image photo format.
func (i *image) Format() string {
	return i.ImageFormat
}

// Equals compares two images for equality. If the two images are the same then true is returned.
// Otherwise, false is returned.
func (i *image) Equals(other Image) bool {
	if i == other {
		return true
	}

	if i.URL() != other.URL() {
		return false
	}

	if i.Size() != other.Size() {
		return false
	}

	if i.Format() != other.Format() {
		return false
	}

	return true
}

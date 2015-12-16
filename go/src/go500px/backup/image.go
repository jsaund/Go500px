package go500px

// JPEG is the JPEG standard photo format.
const JPEG string = "jpeg"

// Image provides the client with the necessary information to retrieve and display an image.
type Image interface {
	URL() string
	Size() Size
	Format() string
}

// ImageImpl implements the Image interface.
type ImageImpl struct {
	ImageURL    string `json:"https_url"`
	SizeID      int    `json:"size"`
	ImageFormat string `json:"format"`
	size        Size
}

type Images struct {
	images []*ImageImpl
}

func NewImages() *Images {
	return new(Images)
}

func (i *Images) Count() int {
	return len(i.images)
}

func (i *Images) Get(index int) *ImageImpl {
	return i.images[index]
}

// URL implements the Image interface. Returns the image URL to download the image.
func (image *ImageImpl) URL() string {
	return image.ImageURL
}

// Size implements the Image interface. Returns the Size of the image in pixels.
func (image *ImageImpl) Size() Size {
	if image.size == nil {
		image.size = NewSizeFromID(image.SizeID)
	}
	return image.size
}

// Format implements the Image interface. Returns the image photo format.
func (image *ImageImpl) Format() string {
	return image.ImageFormat
}

// Equals compares two images for equality. If the two images are the same then true is returned.
// Otherwise, false is returned.
func (image *ImageImpl) Equals(other *ImageImpl) bool {
	if image == other {
		return true
	}

	if image.URL() != other.URL() {
		return false
	}

	if image.Size() != other.Size() {
		return false
	}

	if image.Format() != other.Format() {
		return false
	}

	return true
}

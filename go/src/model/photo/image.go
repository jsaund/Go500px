package photo

// Format defines the image format.
type Format string

// JPEG is the JPEG standard photo format.
const JPEG Format = "jpeg"

// Image provides the client with the necessary information to retrieve and display an image.
type Image interface {
	URL() string
	Size() Size
	Format() Format
}

// ImageImpl implements the Image interface.
type ImageImpl struct {
	ImageURL    string `json:"https_url"`
	SizeID      int    `json:"size"`
	ImageFormat Format `json:"format"`
	size        Size
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
func (image *ImageImpl) Format() Format {
	return image.ImageFormat
}

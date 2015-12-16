package go500px

import "time"

// Photo represents a 500px photo. The photo consists of metadata and
// details regarding retrieving and displaying the image associated with
// this Photo.
type Photo interface {
	ID() int
	Name() string
	Description() string
	Camera() string
	Lens() string
	FocalLength() string
	ISO() string
	ShutterSpeed() string
	Aperture() string
	TimesViewed() int
	Rating() float32
	Status() int
	CreatedAt() int64
	Category() int8
	Location() string
	Privacy() bool
	Latitude() float64
	Longitude() float64
	TakenAt() int64
	ForSale() bool
	Width() int
	Height() int
	VotesCount() int
	FavoritesCount() int
	CommentsCount() int
	PositiveVotesCount() int
	NSFW() bool
	SalesCount() int
	HighestRating() float32
	HighestRatingDate() int64
	LicenseType() int8
	Images() *Images
	User() User
	CollectionsCount() int
}

type photoImpl struct {
	PhotoID                 int          `json:"id"`
	PhotoUser               *UserImpl    `json:"user"`
	PhotoName               string       `json:"name"`
	PhotoDescription        string       `json:"description"`
	PhotoCamera             string       `json:"camera"`
	PhotoLens               string       `json:"lens"`
	PhotoFocalLength        string       `json:"focal_length"`
	PhotoISO                string       `json:"iso"`
	PhotoShutterSpeed       string       `json:"shutter_speed"`
	PhotoAperture           string       `json:"aperture"`
	PhotoTimesViewed        int          `json:"times_viewed"`
	PhotoRating             float32      `json:"rating"`
	PhotoStatus             int          `json:"status"`
	PhotoCreatedAt          *time.Time   `json:"created_at"`
	PhotoCategory           int8         `json:"category"`
	PhotoLocation           string       `json:"location"`
	PhotoLatitude           float64      `json:"latitude"`
	PhotoLongitude          float64      `json:"longitude"`
	PhotoTakenAt            *time.Time   `json:"taken_at"`
	PhotoForSale            bool         `json:"for_sale"`
	PhotoWidth              int          `json:"width"`
	PhotoHeight             int          `json:"height"`
	PhotoVotesCount         int          `json:"votes_count"`
	PhotoFavoritesCount     int          `json:"favorites_count"`
	PhotoCommentsCount      int          `json:"comments_count"`
	PhotoPositiveVotesCount int          `json:"positive_votes_count"`
	PhotoNSFW               bool         `json:"nsfw"`
	PhotoSalesCount         int          `json:"sales_count"`
	PhotoHighestRating      float32      `json:"highest_rating"`
	PhotoHighestRatingDate  *time.Time   `json:"highest_rating_date"`
	PhotoLicenseType        int8         `json:"license_type"`
	PhotoImages             []*ImageImpl `json:"images"`
	PhotoCollectionsCount   int          `json:"collections_count"`
}

func (p *photoImpl) ID() int {
	return p.PhotoID
}

func (p *photoImpl) Name() string {
	return p.PhotoName
}

func (p *photoImpl) Description() string {
	return p.PhotoDescription
}

func (p *photoImpl) Camera() string {
	return p.PhotoCamera
}

func (p *photoImpl) Lens() string {
	return p.PhotoLens
}

func (p *photoImpl) FocalLength() string {
	return p.PhotoFocalLength
}

func (p *photoImpl) ISO() string {
	return p.PhotoISO
}

func (p *photoImpl) ShutterSpeed() string {
	return p.PhotoShutterSpeed
}

func (p *photoImpl) Aperture() string {
	return p.PhotoAperture
}

func (p *photoImpl) TimesViewed() int {
	return p.PhotoTimesViewed
}

func (p *photoImpl) Rating() float32 {
	return p.PhotoRating
}

func (p *photoImpl) Status() int {
	return p.PhotoStatus
}

func (p *photoImpl) CreatedAt() int64 {
	return int64(p.PhotoCreatedAt.Nanosecond() / 1000)
}

func (p *photoImpl) Category() int8 {
	return p.PhotoCategory
}

func (p *photoImpl) Location() string {
	return p.PhotoLocation
}

func (p *photoImpl) Latitude() float64 {
	return p.PhotoLatitude
}

func (p *photoImpl) Longitude() float64 {
	return p.PhotoLongitude
}

func (p *photoImpl) TakenAt() int64 {
	return int64(p.PhotoTakenAt.Nanosecond() / 1000)
}

func (p *photoImpl) ForSale() bool {
	return p.PhotoForSale
}

func (p *photoImpl) Width() int {
	return p.PhotoWidth
}

func (p *photoImpl) Height() int {
	return p.PhotoHeight
}

func (p *photoImpl) VotesCount() int {
	return p.PhotoVotesCount
}

func (p *photoImpl) FavoritesCount() int {
	return p.PhotoFavoritesCount
}

func (p *photoImpl) CommentsCount() int {
	return p.PhotoCommentsCount
}

func (p *photoImpl) PositiveVotesCount() int {
	return p.PhotoPositiveVotesCount
}

func (p *photoImpl) NSFW() bool {
	return p.PhotoNSFW
}

func (p *photoImpl) SalesCount() int {
	return p.PhotoSalesCount
}

func (p *photoImpl) HighestRating() float32 {
	return p.PhotoHighestRating
}

func (p *photoImpl) HighestRatingDate() int64 {
	return int64(p.PhotoHighestRatingDate.Nanosecond() / 1000)
}

func (p *photoImpl) LicenseType() int8 {
	return p.PhotoLicenseType
}

func (p *photoImpl) Images() *Images {
	i := NewImages()
	i.images = p.PhotoImages
	return i
}

func (p *photoImpl) User() *UserImpl {
	return p.PhotoUser
}

func (p *photoImpl) CollectionsCount() int {
	return p.PhotoCollectionsCount
}

func (p *photoImpl) Equals(other *photoImpl) bool {
	if p == other {
		return true
	}

	if other == nil {
		return false
	}

	if p.Aperture() != other.Aperture() {
		return false
	}

	if p.Camera() != other.Camera() {
		return false
	}

	if p.Category() != other.Category() {
		return false
	}

	if p.CollectionsCount() != other.CollectionsCount() {
		return false
	}

	if p.CommentsCount() != other.CommentsCount() {
		return false
	}

	if p.CreatedAt() != other.CreatedAt() {
		return false
	}

	if p.Description() != other.Description() {
		return false
	}

	if p.FavoritesCount() != other.FavoritesCount() {
		return false
	}

	if p.FocalLength() != other.FocalLength() {
		return false
	}

	if p.ForSale() != other.ForSale() {
		return false
	}

	if p.Height() != other.Height() {
		return false
	}

	if p.HighestRating() != other.HighestRating() {
		return false
	}

	if p.HighestRatingDate() != other.HighestRatingDate() {
		return false
	}

	if p.ID() != other.ID() {
		return false
	}

	if p.ISO() != other.ISO() {
		return false
	}

	if p.Images().Count() != other.Images().Count() {
		return false
	}

	for i := 0; i < p.Images().Count(); i++ {
		if !p.Images().Get(i).Equals(other.Images().Get(i)) {
			return false
		}
	}

	if p.Latitude() != other.Latitude() {
		return false
	}

	if p.Longitude() != other.Longitude() {
		return false
	}

	if p.Lens() != other.Lens() {
		return false
	}

	if p.LicenseType() != other.LicenseType() {
		return false
	}

	if p.Location() != other.Location() {
		return false
	}

	if p.NSFW() != other.NSFW() {
		return false
	}

	if p.Name() != other.Name() {
		return false
	}

	if p.PositiveVotesCount() != other.PositiveVotesCount() {
		return false
	}

	if p.Rating() != other.Rating() {
		return false
	}

	if p.SalesCount() != other.SalesCount() {
		return false
	}

	if p.ShutterSpeed() != other.ShutterSpeed() {
		return false
	}

	if p.Status() != other.Status() {
		return false
	}

	if p.TakenAt() != other.TakenAt() {
		return false
	}

	if p.TimesViewed() != other.TimesViewed() {
		return false
	}

	if !p.User().Equals(other.User()) {
		return false
	}

	if p.VotesCount() != other.VotesCount() {
		return false
	}

	if p.Width() != other.Width() {
		return false
	}

	return true
}

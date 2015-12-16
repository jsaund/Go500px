package go500px

import "time"

type Photo interface {
	ID() int32
	Name() string
	Description() string
	Camera() string
	Lens() string
	FocalLength() string
	ISO() string
	ShutterSpeed() string
	Aperture() string
	TimesViewed() int32
	Rating() float32
	Status() int8
	CreatedAt() int64
	Category() int8
	Location() string
	Latitude() float64
	Longitude() float64
	TakenAt() int64
	ForSale() bool
	Width() int32
	Height() int32
	VotesCount() int32
	FavoritesCount() int32
	CommentsCount() int32
	PositiveVotesCount() int32
	NSFW() bool
	SalesCount() int32
	HighestRating() float32
	HighestRatingDate() int64
	LicenseType() int8
	Images() *Images
	User() User
	CollectionsCount() int32
}

type Photos struct {
	photos []*photo
}

func NewPhotos() *Photos {
	return new(Photos)
}

func (p *Photos) Count() int32 {
	return int32(len(p.photos))
}

func (p *Photos) Get(index int32) Photo {
	return p.photos[index]
}

type photo struct {
	PhotoID                 int32      `json:"id"`
	PhotoUser               *user      `json:"user"`
	PhotoName               string     `json:"name"`
	PhotoDescription        string     `json:"description"`
	PhotoCamera             string     `json:"camera"`
	PhotoLens               string     `json:"lens"`
	PhotoFocalLength        string     `json:"focal_length"`
	PhotoISO                string     `json:"iso"`
	PhotoShutterSpeed       string     `json:"shutter_speed"`
	PhotoAperture           string     `json:"aperture"`
	PhotoTimesViewed        int32      `json:"times_viewed"`
	PhotoRating             float32    `json:"rating"`
	PhotoStatus             int8       `json:"status"`
	PhotoCreatedAt          *time.Time `json:"created_at"`
	PhotoCategory           int8       `json:"category"`
	PhotoLocation           string     `json:"location"`
	PhotoLatitude           float64    `json:"latitude"`
	PhotoLongitude          float64    `json:"longitude"`
	PhotoTakenAt            *time.Time `json:"taken_at"`
	PhotoForSale            bool       `json:"for_sale"`
	PhotoWidth              int32      `json:"width"`
	PhotoHeight             int32      `json:"height"`
	PhotoVotesCount         int32      `json:"votes_count"`
	PhotoFavoritesCount     int32      `json:"favorites_count"`
	PhotoCommentsCount      int32      `json:"comments_count"`
	PhotoPositiveVotesCount int32      `json:"positive_votes_count"`
	PhotoNSFW               bool       `json:"nsfw"`
	PhotoSalesCount         int32      `json:"sales_count"`
	PhotoHighestRating      float32    `json:"highest_rating"`
	PhotoHighestRatingDate  *time.Time `json:"highest_rating_date"`
	PhotoLicenseType        int8       `json:"license_type"`
	PhotoImages             []*image   `json:"images"`
	PhotoCollectionsCount   int32      `json:"collections_count"`
}

func (p *photo) ID() int32 {
	return p.PhotoID
}

func (p *photo) Name() string {
	return p.PhotoName
}

func (p *photo) Description() string {
	return p.PhotoDescription
}

func (p *photo) Camera() string {
	return p.PhotoCamera
}

func (p *photo) Lens() string {
	return p.PhotoLens
}

func (p *photo) FocalLength() string {
	return p.PhotoFocalLength
}

func (p *photo) ISO() string {
	return p.PhotoISO
}

func (p *photo) ShutterSpeed() string {
	return p.PhotoShutterSpeed
}

func (p *photo) Aperture() string {
	return p.PhotoAperture
}

func (p *photo) TimesViewed() int32 {
	return p.PhotoTimesViewed
}

func (p *photo) Rating() float32 {
	return p.PhotoRating
}

func (p *photo) Status() int8 {
	return p.PhotoStatus
}

func (p *photo) CreatedAt() int64 {
	return int64(p.PhotoCreatedAt.Nanosecond() / 1000)
}

func (p *photo) Category() int8 {
	return p.PhotoCategory
}

func (p *photo) Location() string {
	return p.PhotoLocation
}

func (p *photo) Latitude() float64 {
	return p.PhotoLatitude
}

func (p *photo) Longitude() float64 {
	return p.PhotoLongitude
}

func (p *photo) TakenAt() int64 {
	return int64(p.PhotoTakenAt.Nanosecond() / 1000)
}

func (p *photo) ForSale() bool {
	return p.PhotoForSale
}

func (p *photo) Width() int32 {
	return p.PhotoWidth
}

func (p *photo) Height() int32 {
	return p.PhotoHeight
}

func (p *photo) VotesCount() int32 {
	return p.PhotoVotesCount
}

func (p *photo) FavoritesCount() int32 {
	return p.PhotoFavoritesCount
}

func (p *photo) CommentsCount() int32 {
	return p.PhotoCommentsCount
}

func (p *photo) PositiveVotesCount() int32 {
	return p.PhotoPositiveVotesCount
}

func (p *photo) NSFW() bool {
	return p.PhotoNSFW
}

func (p *photo) SalesCount() int32 {
	return p.PhotoSalesCount
}

func (p *photo) HighestRating() float32 {
	return p.PhotoHighestRating
}

func (p *photo) HighestRatingDate() int64 {
	return int64(p.PhotoHighestRatingDate.Nanosecond() / 1000)
}

func (p *photo) LicenseType() int8 {
	return p.PhotoLicenseType
}

func (p *photo) Images() *Images {
	i := NewImages()
	i.images = p.PhotoImages
	return i
}

func (p *photo) User() User {
	return p.PhotoUser
}

func (p *photo) CollectionsCount() int32 {
	return p.PhotoCollectionsCount
}

func (p *photo) Equals(other *photo) bool {
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

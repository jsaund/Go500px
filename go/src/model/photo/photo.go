package photo

import "time"

import "model/category"
import "model/user"
import "model/license"

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
	Category() category.Category
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
	LicenseType() license.LicenseType
	Images() []Image
	User() user.User
	CollectionsCount() int
}

type PhotoImpl struct {
	PhotoID                 int                 `json:"id"`
	PhotoUser               *user.UserImpl      `json:"user"`
	PhotoName               string              `json:"name"`
	PhotoDescription        string              `json:"description"`
	PhotoCamera             string              `json:"camera"`
	PhotoLens               string              `json:"lens"`
	PhotoISO                string              `json:"iso"`
	PhotoShutterSpeed       string              `json:"shutter_speed"`
	PhotoAperture           string              `json:"aperture"`
	PhotoTimesViewed        int                 `json:"times_viewed"`
	PhotoRating             float32             `json:"rating"`
	PhotoStatus             int                 `json:"status"`
	PhotoCreatedAt          time.Time           `json:"created_at"`
	PhotoCategory           category.Category   `json:"category"`
	PhotoLocation           string              `json:"location"`
	PhotoLatitude           float64             `json:"latitude"`
	PhotoLongitude          float64             `json:"longitude"`
	PhotoTakenAt            time.Time           `json:"taken_at"`
	PhotoForSale            bool                `json:"for_sale"`
	PhotoWidth              int                 `json:"width"`
	PhotoHeight             int                 `json:"height"`
	PhotoVotesCount         int                 `json:"votes_count"`
	PhotoFavoritesCount     int                 `json:"favorites_count"`
	PhotoCommentsCount      int                 `json:"comments_count"`
	PhotoPositiveVotesCount int                 `json:"positive_votes_count"`
	PhotoNSFW               bool                `json:"nsfw"`
	PhotoSalesCount         int                 `json:"sales_count"`
	PhotoHighestRating      float32             `json:"highest_rating"`
	PhotoHighestRatingDate  time.Time           `json:"highest_rating_date"`
	PhotoLicenseType        license.LicenseType `json:"license_type"`
	PhotoImages             []*ImageImpl        `json:"images"`
	PhotoCollectionsCount   int                 `json:"collections_count"`
}

func (p *PhotoImpl) ID() int {
	return p.PhotoID
}

func (p *PhotoImpl) Name() string {
	return p.PhotoName
}

func (p *PhotoImpl) Description() string {
	return p.PhotoDescription
}

func (p *PhotoImpl) Camera() string {
	return p.PhotoCamera
}

func (p *PhotoImpl) Lens() string {
	return p.PhotoLens
}

func (p *PhotoImpl) ISO() string {
	return p.PhotoISO
}

func (p *PhotoImpl) ShutterSpeed() string {
	return p.PhotoShutterSpeed
}

func (p *PhotoImpl) Aperture() string {
	return p.PhotoAperture
}

func (p *PhotoImpl) TimesViewed() int {
	return p.PhotoTimesViewed
}

func (p *PhotoImpl) Rating() float32 {
	return p.PhotoRating
}

func (p *PhotoImpl) Status() int {
	return p.PhotoStatus
}

func (p *PhotoImpl) CreatedAt() int64 {
	return int64(p.PhotoCreatedAt.Nanosecond() / 1000)
}

func (p *PhotoImpl) Category() category.Category {
	return p.PhotoCategory
}

func (p *PhotoImpl) Location() string {
	return p.PhotoLocation
}

func (p *PhotoImpl) Latitude() float64 {
	return p.PhotoLatitude
}

func (p *PhotoImpl) Longitude() float64 {
	return p.PhotoLongitude
}

func (p *PhotoImpl) TakenAt() int64 {
	return int64(p.PhotoTakenAt.Nanosecond() / 1000)
}

func (p *PhotoImpl) ForSale() bool {
	return p.PhotoForSale
}

func (p *PhotoImpl) Width() int {
	return p.PhotoWidth
}

func (p *PhotoImpl) Height() int {
	return p.PhotoHeight
}

func (p *PhotoImpl) VotesCount() int {
	return p.PhotoVotesCount
}

func (p *PhotoImpl) FavoritesCount() int {
	return p.PhotoFavoritesCount
}

func (p *PhotoImpl) CommentsCount() int {
	return p.PhotoCommentsCount
}

func (p *PhotoImpl) PositiveVotesCount() int {
	return p.PhotoPositiveVotesCount
}

func (p *PhotoImpl) NSFW() bool {
	return p.PhotoNSFW
}

func (p *PhotoImpl) SalesCount() int {
	return p.PhotoSalesCount
}

func (p *PhotoImpl) HighestRating() float32 {
	return p.PhotoHighestRating
}

func (p *PhotoImpl) HighestRatingDate() int64 {
	return int64(p.PhotoHighestRatingDate.Nanosecond() / 1000)
}

func (p *PhotoImpl) LicenseType() license.LicenseType {
	return p.PhotoLicenseType
}

func (p *PhotoImpl) Images() []*ImageImpl {
	return p.PhotoImages
}

func (p *PhotoImpl) User() user.User {
	return p.PhotoUser
}

func (p *PhotoImpl) CollectionsCount() int {
	return p.PhotoCollectionsCount
}
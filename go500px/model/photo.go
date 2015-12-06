package model

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
	Category() Category
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
	LicenseType() LicenseType
	Images() []Image
	User() User
	CollectionsCount() int
}

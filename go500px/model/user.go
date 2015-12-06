package model

type User interface {
	ID() int
	Username() string
	Firstname() string
	Lastname() string
	Fullname() string
	City() string
	Country() string
	UserPicURL() string
	UpgradeStatus() int
	FollowersCount() int
	Affection() int
}

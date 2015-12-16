package go500px

// User defines user metadata information.
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
	Equals(other User) bool
}

// user implements the User interface.
type user struct {
	UserID             int    `json:"id"`
	UserUsername       string `json:"username"`
	UserFirstname      string `json:"firstname"`
	UserLastname       string `json:"lastname"`
	UserFullname       string `json:"fullname"`
	UserCity           string `json:"city"`
	UserCountry        string `json:"country"`
	UserUserPicURL     string `json:"userpic_https_url"`
	UserUpgradeStatus  int    `json:"upgrade_status"`
	UserFollowersCount int    `json:"followers_count"`
	UserAffection      int    `json:"affection"`
}

// ID returns the user's ID
func (u *user) ID() int {
	return u.UserID
}

// Username returns the user's username
func (u *user) Username() string {
	return u.UserUsername
}

// Firstname returns the user's first name
func (u *user) Firstname() string {
	return u.UserFirstname
}

// Lastname returns the user's last name
func (u *user) Lastname() string {
	return u.UserLastname
}

// Fullname returns the user's full name as a concatenation of the
// user's first name and last name.
func (u *user) Fullname() string {
	return u.UserFullname
}

// City returns the user's city
func (u *user) City() string {
	return u.UserCity
}

// Country returns the user's country
func (u *user) Country() string {
	return u.UserCountry
}

// UserPicURL returns the url for the user's profile picture
func (u *user) UserPicURL() string {
	return u.UserUserPicURL
}

// UpgradeStatus returns whether the user is a premium user. Non-zero values identify premium users;
// a value of 2 identifies an Awesome user while a value of 1 identifies a Plus user.
// Other states may be added in the future.
func (u *user) UpgradeStatus() int {
	return u.UserUpgradeStatus
}

// FollowersCount returns the user's followers count
func (u *user) FollowersCount() int {
	return u.UserFollowersCount
}

// Affection returns the user's affection value
func (u *user) Affection() int {
	return u.UserAffection
}

// Equals compares the equality of two users. If the two users are the same, then true is returned. Otherwise
// false is returned.
func (u *user) Equals(other User) bool {
	if other == nil {
		return false
	}

	if u.Affection() != other.Affection() {
		return false
	}

	if u.City() != other.City() {
		return false
	}

	if u.Country() != other.Country() {
		return false
	}

	if u.Firstname() != other.Firstname() {
		return false
	}

	if u.FollowersCount() != other.FollowersCount() {
		return false
	}

	if u.Fullname() != other.Fullname() {
		return false
	}

	if u.ID() != other.ID() {
		return false
	}

	if u.Lastname() != other.Lastname() {
		return false
	}

	if u.UpgradeStatus() != other.UpgradeStatus() {
		return false
	}

	if u.Username() != other.Username() {
		return false
	}

	if u.UserPicURL() != other.UserPicURL() {
		return false
	}

	return true
}

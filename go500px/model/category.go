package model

// Category defines the type of category a Photo belongs to.
type Category int8

const (
	// Uncategorized category
	Uncategorized Category = iota
	// Celebrities category
	Celebrities
	// Film category
	Film
	// Journalism category
	Journalism
	// Nude category
	Nude
	// BlackAndWhite category
	BlackAndWhite
	// StillLife category
	StillLife
	// People category
	People
	// Landscapes category
	Landscapes
	// CityAndArchitecture category
	CityAndArchitecture
	// Abstract category
	Abstract
	// Animals category
	Animals
	// Macro category
	Macro
	// Travel category
	Travel
	// Fashion category
	Fashion
	// Commercial category
	Commercial
	// Concert category
	Concert
	// Sport category
	Sport
	// Nature category
	Nature
	// PerformingArts category
	PerformingArts
	// Family category
	Family
	// Street category
	Street
	// Underwater category
	Underwater
	// Food category
	Food
	// FineArt category
	FineArt
	// Wedding category
	Wedding
	// Transportation category
	Transportation
	// UrbanExploration category
	UrbanExploration
)

categoryToString := map[Category]string {
}

func NewCategoryFromString(category string) Category {

}

func (c Category) String() string {
	switch {
	case Uncategorized:
		return "Uncategorized"
	case Celebrities:
		return "Celebrities"
	case Film:
		return "Film"
	case Journalism:
		return "Journalism"
	case Nude:
		return "Nude"
	case BlackAndWhite:
		return "Black and White"
	case StillLife:
		return "Still Life"
	case People:
		return "People"
	case Landscapes:
		return "Landscapes"
	case CityAndArchitecture:
		return "City and Architecture"
	case Abstract:
		return "Abstract"
	case Macro:
		return "Macro"
	case Travel:
		return "Travel"
	case Fashion:
		return "Fashion"
	case Commercial:
		return "Commercial"
	case Concert:
		return "Concert"
	case Sport:
		return "Sport"
	case Nature:
		return "Nature"
	case PerformingArts:
		return "Performing Arts"
	case Family:
		return "Family"
	case Street:
		return "Street"
	case Underwater:
		return "Underwater"
	case Food:
		return "Food"
	case FineArt:
		return "Fine Art"
	case Wedding:
		return "Wedding"
	case Transportation:
		return "Transportation"
	case UrbanExploration:
		return "Urban Exploration"
	default:
		return ""
	}
}

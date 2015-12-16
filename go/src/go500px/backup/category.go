package go500px

// Category defines the type of category a Photo belongs to.
const (
	// Uncategorized category
	Uncategorized int8 = iota
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

var categoryToString = map[int8]string{
	Uncategorized:       "Uncategorized",
	Celebrities:         "Celebrities",
	Film:                "Film",
	Journalism:          "Journalism",
	Nude:                "Nude",
	BlackAndWhite:       "Black and White",
	StillLife:           "Still Life",
	People:              "People",
	Landscapes:          "Landscapes",
	CityAndArchitecture: "City and Architecture",
	Abstract:            "Abstract",
	Macro:               "Macro",
	Travel:              "Travel",
	Fashion:             "Fashion",
	Commercial:          "Commercial",
	Concert:             "Concert",
	Sport:               "Sport",
	Nature:              "Nature",
	PerformingArts:      "Performing Arts",
	Family:              "Family",
	Street:              "Street",
	Underwater:          "Underwater",
	Food:                "Food",
	FineArt:             "FineArt",
	Wedding:             "Wedding",
	Transportation:      "Transportation",
	UrbanExploration:    "Urban Exploration",
}

// NewCategoryFromString returns the Category which corresponds to the literal category value.
func NewCategoryFromString(category string) int8 {
	for c, s := range categoryToString {
		if s == category {
			return c
		}
	}
	return Uncategorized
}

func CategoryToString(c int8) string {
	s, ok := categoryToString[c]
	if ok {
		return s
	}
	return categoryToString[Uncategorized]
}

package go500px

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"
)

type photoTest struct {
	Photo *photoImpl `json:"photo"`
}

func TestParsePhoto(t *testing.T) {
	filename := "photo_test_1.json"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to open file: %s. Reason: %v", filename, err)
	}

	createdAt, _ := time.Parse(time.RFC3339, "2015-12-07T03:26:06-05:00")
	takenAt, _ := time.Parse(time.RFC3339, "2015-08-08T05:38:39-04:00")
	highestRatingDate, _ := time.Parse(time.RFC3339, "2015-12-07T18:29:54-05:00")

	images := []*ImageImpl{
		&ImageImpl{
			ImageFormat: JPEG,
			SizeID:      3,
			ImageURL:    "https://drscdn.500px.org/photo/131478301/w%3D280_h%3D280/a6ee6d7e1e291a4622e1679d38e61a5a?v=3",
		},
	}

	user := &UserImpl{
		UserID:             1568405,
		UserUsername:       "ZekiSeferoglu",
		UserFirstname:      "Zeki",
		UserLastname:       "Seferoglu",
		UserFullname:       "Zeki Seferoglu",
		UserCity:           "AKSARAY",
		UserCountry:        "TURKEY",
		UserUserPicURL:     "https://pacdn.500px.org/1568405/89d1765e4bdb376ac4647358b7a77ca225970cb7/1.jpg?9",
		UserUpgradeStatus:  0,
		UserFollowersCount: 2269,
		UserAffection:      86791,
	}

	expected1 := &photoImpl{
		PhotoID:                 131478301,
		PhotoUser:               user,
		PhotoName:               "Perilerle Gelen Sabah",
		PhotoDescription:        "G\u00f6reme / Turkey",
		PhotoCamera:             "NIKON D810",
		PhotoLens:               "70.0-200.0 mm f/2.8",
		PhotoFocalLength:        "135",
		PhotoISO:                "100",
		PhotoShutterSpeed:       "1/2500",
		PhotoAperture:           "2.8",
		PhotoTimesViewed:        5213,
		PhotoRating:             99.5,
		PhotoStatus:             1,
		PhotoCreatedAt:          &createdAt,
		PhotoCategory:           Landscapes,
		PhotoLocation:           "",
		PhotoLatitude:           38.6528070477738,
		PhotoLongitude:          34.8362731933594,
		PhotoTakenAt:            &takenAt,
		PhotoForSale:            false,
		PhotoWidth:              1920,
		PhotoHeight:             1281,
		PhotoVotesCount:         1046,
		PhotoFavoritesCount:     106,
		PhotoCommentsCount:      109,
		PhotoPositiveVotesCount: 1046,
		PhotoNSFW:               false,
		PhotoSalesCount:         0,
		PhotoHighestRating:      99.5,
		PhotoHighestRatingDate:  &highestRatingDate,
		PhotoLicenseType:        Standard,
		PhotoImages:             images,
		PhotoCollectionsCount:   110,
	}

	var result1 photoTest
	if json.Unmarshal(bytes, &result1) != nil {
		t.Fatalf("Failed to unmarshal json. Reason: %v", err)
	}

	if result1.Photo == nil {
		t.Fatal("Failed to unmarshal json. Photo is nil")
	}

	if !result1.Photo.Equals(expected1) {
		t.Errorf("Unexpected values received in unmarshaled photo. Result: %v\n\nExpcted: %v",
			result1.Photo, expected1)
	}
}

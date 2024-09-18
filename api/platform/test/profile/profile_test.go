package profile_test

import (
	"spc/api/platform/src/profile"
	"testing"
)

func TestAdd(t *testing.T) {
	collection := profile.New()
	collection.Add(profile.Album{})
	if len(collection.Albums) != 1 {
		t.Errorf("Album was not added to Collection")
	}
}

func TestGetAll(t *testing.T) {
	collection := profile.New()
	collection.Add(profile.Album{})
	result := collection.GetAll()
	if len(result) != 1 {
		t.Errorf("Album was not added to Collection")
	}
}

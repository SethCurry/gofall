package gofall

import "errors"

// ErrNoImageURIs is returned when a card has no image URIs.
var ErrNoImageURIs = errors.New("no image URIs for card")

type ImageURIs struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	PNG        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

// LowestQuality returns the lowest quality image URI for the card.
// Prefers jpeg images.
func (i *ImageURIs) LowestQuality() (string, error) {
	if i.Small != "" {
		return i.Small, nil
	}

	if i.Normal != "" {
		return i.Normal, nil
	}

	if i.Large != "" {
		return i.Large, nil
	}

	if i.PNG != "" {
		return i.PNG, nil
	}

	if i.ArtCrop != "" {
		return i.ArtCrop, nil
	}

	if i.BorderCrop != "" {
		return i.BorderCrop, nil
	}

	return "", ErrNoImageURIs
}

// HighestQuality returns the highest quality image URI for the card.
// Prefers jpeg images.
func (i *ImageURIs) HighestQuality() (string, error) {
	if i.Large != "" {
		return i.Large, nil
	}

	if i.Normal != "" {
		return i.Normal, nil
	}

	if i.Small != "" {
		return i.Small, nil
	}

	if i.PNG != "" {
		return i.PNG, nil
	}

	if i.ArtCrop != "" {
		return i.ArtCrop, nil
	}

	if i.BorderCrop != "" {
		return i.BorderCrop, nil
	}

	return "", ErrNoImageURIs
}

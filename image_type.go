package gofall

// ImageType defines the types of images such as small, normal, large, etc.
type ImageType string

const (
	// ImageTypeSmall is the smallest image size.
	ImageTypeSmall ImageType = "small"

	// ImageTypeNormal is the normal image size.
	ImageTypeNormal ImageType = "normal"

	// ImageTypeLarge is the largest image size.
	ImageTypeLarge ImageType = "large"

	// ImageTypePng is the PNG image type.
	ImageTypePng ImageType = "png"

	// ImageTypeArtCrop is the art crop image type.
	ImageTypeArtCrop ImageType = "art_crop"

	// ImageTypeBorderCrop is the border crop image type.
	ImageTypeBorderCrop ImageType = "border_crop"
)

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RatingJapanTelevisionType undocumented
type RatingJapanTelevisionType int

const (
	// RatingJapanTelevisionTypeVAllAllowed undocumented
	RatingJapanTelevisionTypeVAllAllowed RatingJapanTelevisionType = 0
	// RatingJapanTelevisionTypeVAllBlocked undocumented
	RatingJapanTelevisionTypeVAllBlocked RatingJapanTelevisionType = 1
	// RatingJapanTelevisionTypeVExplicitAllowed undocumented
	RatingJapanTelevisionTypeVExplicitAllowed RatingJapanTelevisionType = 2
)

// RatingJapanTelevisionTypePAllAllowed returns a pointer to RatingJapanTelevisionTypeVAllAllowed
func RatingJapanTelevisionTypePAllAllowed() *RatingJapanTelevisionType {
	v := RatingJapanTelevisionTypeVAllAllowed
	return &v
}

// RatingJapanTelevisionTypePAllBlocked returns a pointer to RatingJapanTelevisionTypeVAllBlocked
func RatingJapanTelevisionTypePAllBlocked() *RatingJapanTelevisionType {
	v := RatingJapanTelevisionTypeVAllBlocked
	return &v
}

// RatingJapanTelevisionTypePExplicitAllowed returns a pointer to RatingJapanTelevisionTypeVExplicitAllowed
func RatingJapanTelevisionTypePExplicitAllowed() *RatingJapanTelevisionType {
	v := RatingJapanTelevisionTypeVExplicitAllowed
	return &v
}
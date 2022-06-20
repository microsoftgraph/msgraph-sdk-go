package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaContentRatingUnitedKingdom 
type MediaContentRatingUnitedKingdom struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Movies rating selected for United Kingdom. Possible values are: allAllowed, allBlocked, general, universalChildren, parentalGuidance, agesAbove12Video, agesAbove12Cinema, agesAbove15, adults.
    movieRating *RatingUnitedKingdomMoviesType
    // TV rating selected for United Kingdom. Possible values are: allAllowed, allBlocked, caution.
    tvRating *RatingUnitedKingdomTelevisionType
}
// NewMediaContentRatingUnitedKingdom instantiates a new mediaContentRatingUnitedKingdom and sets the default values.
func NewMediaContentRatingUnitedKingdom()(*MediaContentRatingUnitedKingdom) {
    m := &MediaContentRatingUnitedKingdom{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaContentRatingUnitedKingdomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaContentRatingUnitedKingdomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaContentRatingUnitedKingdom(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaContentRatingUnitedKingdom) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaContentRatingUnitedKingdom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["movieRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingUnitedKingdomMoviesType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMovieRating(val.(*RatingUnitedKingdomMoviesType))
        }
        return nil
    }
    res["tvRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingUnitedKingdomTelevisionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTvRating(val.(*RatingUnitedKingdomTelevisionType))
        }
        return nil
    }
    return res
}
// GetMovieRating gets the movieRating property value. Movies rating selected for United Kingdom. Possible values are: allAllowed, allBlocked, general, universalChildren, parentalGuidance, agesAbove12Video, agesAbove12Cinema, agesAbove15, adults.
func (m *MediaContentRatingUnitedKingdom) GetMovieRating()(*RatingUnitedKingdomMoviesType) {
    if m == nil {
        return nil
    } else {
        return m.movieRating
    }
}
// GetTvRating gets the tvRating property value. TV rating selected for United Kingdom. Possible values are: allAllowed, allBlocked, caution.
func (m *MediaContentRatingUnitedKingdom) GetTvRating()(*RatingUnitedKingdomTelevisionType) {
    if m == nil {
        return nil
    } else {
        return m.tvRating
    }
}
// Serialize serializes information the current object
func (m *MediaContentRatingUnitedKingdom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMovieRating() != nil {
        cast := (*m.GetMovieRating()).String()
        err := writer.WriteStringValue("movieRating", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTvRating() != nil {
        cast := (*m.GetTvRating()).String()
        err := writer.WriteStringValue("tvRating", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaContentRatingUnitedKingdom) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMovieRating sets the movieRating property value. Movies rating selected for United Kingdom. Possible values are: allAllowed, allBlocked, general, universalChildren, parentalGuidance, agesAbove12Video, agesAbove12Cinema, agesAbove15, adults.
func (m *MediaContentRatingUnitedKingdom) SetMovieRating(value *RatingUnitedKingdomMoviesType)() {
    if m != nil {
        m.movieRating = value
    }
}
// SetTvRating sets the tvRating property value. TV rating selected for United Kingdom. Possible values are: allAllowed, allBlocked, caution.
func (m *MediaContentRatingUnitedKingdom) SetTvRating(value *RatingUnitedKingdomTelevisionType)() {
    if m != nil {
        m.tvRating = value
    }
}

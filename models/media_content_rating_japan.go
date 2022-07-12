package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaContentRatingJapan 
type MediaContentRatingJapan struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Movies rating labels in Japan
    movieRating *RatingJapanMoviesType
    // TV content rating labels in Japan
    tvRating *RatingJapanTelevisionType
}
// NewMediaContentRatingJapan instantiates a new mediaContentRatingJapan and sets the default values.
func NewMediaContentRatingJapan()(*MediaContentRatingJapan) {
    m := &MediaContentRatingJapan{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaContentRatingJapanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaContentRatingJapanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaContentRatingJapan(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaContentRatingJapan) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaContentRatingJapan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["movieRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingJapanMoviesType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMovieRating(val.(*RatingJapanMoviesType))
        }
        return nil
    }
    res["tvRating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingJapanTelevisionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTvRating(val.(*RatingJapanTelevisionType))
        }
        return nil
    }
    return res
}
// GetMovieRating gets the movieRating property value. Movies rating labels in Japan
func (m *MediaContentRatingJapan) GetMovieRating()(*RatingJapanMoviesType) {
    if m == nil {
        return nil
    } else {
        return m.movieRating
    }
}
// GetTvRating gets the tvRating property value. TV content rating labels in Japan
func (m *MediaContentRatingJapan) GetTvRating()(*RatingJapanTelevisionType) {
    if m == nil {
        return nil
    } else {
        return m.tvRating
    }
}
// Serialize serializes information the current object
func (m *MediaContentRatingJapan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MediaContentRatingJapan) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMovieRating sets the movieRating property value. Movies rating labels in Japan
func (m *MediaContentRatingJapan) SetMovieRating(value *RatingJapanMoviesType)() {
    if m != nil {
        m.movieRating = value
    }
}
// SetTvRating sets the tvRating property value. TV content rating labels in Japan
func (m *MediaContentRatingJapan) SetTvRating(value *RatingJapanTelevisionType)() {
    if m != nil {
        m.tvRating = value
    }
}

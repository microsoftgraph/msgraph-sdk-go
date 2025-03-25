package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceAnnouncementMessagesFavoritepostResponseable instead.
type ServiceAnnouncementMessagesFavoriteResponse struct {
    ServiceAnnouncementMessagesFavoritepostResponse
}
// NewServiceAnnouncementMessagesFavoriteResponse instantiates a new ServiceAnnouncementMessagesFavoriteResponse and sets the default values.
func NewServiceAnnouncementMessagesFavoriteResponse()(*ServiceAnnouncementMessagesFavoriteResponse) {
    m := &ServiceAnnouncementMessagesFavoriteResponse{
        ServiceAnnouncementMessagesFavoritepostResponse: *NewServiceAnnouncementMessagesFavoritepostResponse(),
    }
    return m
}
// CreateServiceAnnouncementMessagesFavoriteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceAnnouncementMessagesFavoriteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementMessagesFavoriteResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceAnnouncementMessagesFavoritepostResponseable instead.
type ServiceAnnouncementMessagesFavoriteResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAnnouncementMessagesFavoritepostResponseable
}

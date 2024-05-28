package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceannouncementMessagesFavoritePostResponseable instead.
type ServiceannouncementMessagesFavoriteResponse struct {
    ServiceannouncementMessagesFavoritePostResponse
}
// NewServiceannouncementMessagesFavoriteResponse instantiates a new ServiceannouncementMessagesFavoriteResponse and sets the default values.
func NewServiceannouncementMessagesFavoriteResponse()(*ServiceannouncementMessagesFavoriteResponse) {
    m := &ServiceannouncementMessagesFavoriteResponse{
        ServiceannouncementMessagesFavoritePostResponse: *NewServiceannouncementMessagesFavoritePostResponse(),
    }
    return m
}
// CreateServiceannouncementMessagesFavoriteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceannouncementMessagesFavoriteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceannouncementMessagesFavoriteResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceannouncementMessagesFavoritePostResponseable instead.
type ServiceannouncementMessagesFavoriteResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceannouncementMessagesFavoritePostResponseable
}

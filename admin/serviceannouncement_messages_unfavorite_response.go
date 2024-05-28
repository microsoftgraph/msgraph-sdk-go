package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceannouncementMessagesUnfavoritePostResponseable instead.
type ServiceannouncementMessagesUnfavoriteResponse struct {
    ServiceannouncementMessagesUnfavoritePostResponse
}
// NewServiceannouncementMessagesUnfavoriteResponse instantiates a new ServiceannouncementMessagesUnfavoriteResponse and sets the default values.
func NewServiceannouncementMessagesUnfavoriteResponse()(*ServiceannouncementMessagesUnfavoriteResponse) {
    m := &ServiceannouncementMessagesUnfavoriteResponse{
        ServiceannouncementMessagesUnfavoritePostResponse: *NewServiceannouncementMessagesUnfavoritePostResponse(),
    }
    return m
}
// CreateServiceannouncementMessagesUnfavoriteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceannouncementMessagesUnfavoriteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceannouncementMessagesUnfavoriteResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceannouncementMessagesUnfavoritePostResponseable instead.
type ServiceannouncementMessagesUnfavoriteResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceannouncementMessagesUnfavoritePostResponseable
}

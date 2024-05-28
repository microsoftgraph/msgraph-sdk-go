package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServiceannouncementMessagesArchivePostResponseable instead.
type ServiceannouncementMessagesArchiveResponse struct {
    ServiceannouncementMessagesArchivePostResponse
}
// NewServiceannouncementMessagesArchiveResponse instantiates a new ServiceannouncementMessagesArchiveResponse and sets the default values.
func NewServiceannouncementMessagesArchiveResponse()(*ServiceannouncementMessagesArchiveResponse) {
    m := &ServiceannouncementMessagesArchiveResponse{
        ServiceannouncementMessagesArchivePostResponse: *NewServiceannouncementMessagesArchivePostResponse(),
    }
    return m
}
// CreateServiceannouncementMessagesArchiveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceannouncementMessagesArchiveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceannouncementMessagesArchiveResponse(), nil
}
// Deprecated: This class is obsolete. Use ServiceannouncementMessagesArchivePostResponseable instead.
type ServiceannouncementMessagesArchiveResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceannouncementMessagesArchivePostResponseable
}

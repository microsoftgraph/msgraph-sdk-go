package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementMessagesUnarchiveResponse 
// Deprecated: This class is obsolete. Use unarchivePostResponse instead.
type ServiceAnnouncementMessagesUnarchiveResponse struct {
    ServiceAnnouncementMessagesUnarchivePostResponse
}
// NewServiceAnnouncementMessagesUnarchiveResponse instantiates a new ServiceAnnouncementMessagesUnarchiveResponse and sets the default values.
func NewServiceAnnouncementMessagesUnarchiveResponse()(*ServiceAnnouncementMessagesUnarchiveResponse) {
    m := &ServiceAnnouncementMessagesUnarchiveResponse{
        ServiceAnnouncementMessagesUnarchivePostResponse: *NewServiceAnnouncementMessagesUnarchivePostResponse(),
    }
    return m
}
// CreateServiceAnnouncementMessagesUnarchiveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementMessagesUnarchiveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementMessagesUnarchiveResponse(), nil
}
// ServiceAnnouncementMessagesUnarchiveResponseable 
// Deprecated: This class is obsolete. Use unarchivePostResponse instead.
type ServiceAnnouncementMessagesUnarchiveResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAnnouncementMessagesUnarchivePostResponseable
}

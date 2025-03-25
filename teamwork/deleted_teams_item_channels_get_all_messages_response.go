package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsGetAllMessagesgetResponseable instead.
type DeletedTeamsItemChannelsGetAllMessagesResponse struct {
    DeletedTeamsItemChannelsGetAllMessagesgetResponse
}
// NewDeletedTeamsItemChannelsGetAllMessagesResponse instantiates a new DeletedTeamsItemChannelsGetAllMessagesResponse and sets the default values.
func NewDeletedTeamsItemChannelsGetAllMessagesResponse()(*DeletedTeamsItemChannelsGetAllMessagesResponse) {
    m := &DeletedTeamsItemChannelsGetAllMessagesResponse{
        DeletedTeamsItemChannelsGetAllMessagesgetResponse: *NewDeletedTeamsItemChannelsGetAllMessagesgetResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsGetAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsGetAllMessagesgetResponseable instead.
type DeletedTeamsItemChannelsGetAllMessagesResponseable interface {
    DeletedTeamsItemChannelsGetAllMessagesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemAllMembersRemovepostResponseable instead.
type DeletedTeamsItemChannelsItemAllMembersRemoveResponse struct {
    DeletedTeamsItemChannelsItemAllMembersRemovepostResponse
}
// NewDeletedTeamsItemChannelsItemAllMembersRemoveResponse instantiates a new DeletedTeamsItemChannelsItemAllMembersRemoveResponse and sets the default values.
func NewDeletedTeamsItemChannelsItemAllMembersRemoveResponse()(*DeletedTeamsItemChannelsItemAllMembersRemoveResponse) {
    m := &DeletedTeamsItemChannelsItemAllMembersRemoveResponse{
        DeletedTeamsItemChannelsItemAllMembersRemovepostResponse: *NewDeletedTeamsItemChannelsItemAllMembersRemovepostResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsItemAllMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsItemAllMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsItemAllMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemAllMembersRemovepostResponseable instead.
type DeletedTeamsItemChannelsItemAllMembersRemoveResponseable interface {
    DeletedTeamsItemChannelsItemAllMembersRemovepostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsItemChannelsGetAllMessagesResponse 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemJoinedTeamsItemChannelsGetAllMessagesResponse struct {
    ItemJoinedTeamsItemChannelsGetAllMessagesGetResponse
}
// NewItemJoinedTeamsItemChannelsGetAllMessagesResponse instantiates a new ItemJoinedTeamsItemChannelsGetAllMessagesResponse and sets the default values.
func NewItemJoinedTeamsItemChannelsGetAllMessagesResponse()(*ItemJoinedTeamsItemChannelsGetAllMessagesResponse) {
    m := &ItemJoinedTeamsItemChannelsGetAllMessagesResponse{
        ItemJoinedTeamsItemChannelsGetAllMessagesGetResponse: *NewItemJoinedTeamsItemChannelsGetAllMessagesGetResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemChannelsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsItemChannelsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemChannelsGetAllMessagesResponse(), nil
}
// ItemJoinedTeamsItemChannelsGetAllMessagesResponseable 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemJoinedTeamsItemChannelsGetAllMessagesResponseable interface {
    ItemJoinedTeamsItemChannelsGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

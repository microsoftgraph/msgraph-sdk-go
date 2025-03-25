package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsGetAllMessagesgetResponseable instead.
type ItemTeamChannelsGetAllMessagesResponse struct {
    ItemTeamChannelsGetAllMessagesgetResponse
}
// NewItemTeamChannelsGetAllMessagesResponse instantiates a new ItemTeamChannelsGetAllMessagesResponse and sets the default values.
func NewItemTeamChannelsGetAllMessagesResponse()(*ItemTeamChannelsGetAllMessagesResponse) {
    m := &ItemTeamChannelsGetAllMessagesResponse{
        ItemTeamChannelsGetAllMessagesgetResponse: *NewItemTeamChannelsGetAllMessagesgetResponse(),
    }
    return m
}
// CreateItemTeamChannelsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsGetAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsGetAllMessagesgetResponseable instead.
type ItemTeamChannelsGetAllMessagesResponseable interface {
    ItemTeamChannelsGetAllMessagesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

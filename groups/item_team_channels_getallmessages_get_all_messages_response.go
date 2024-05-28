package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsGetallmessagesGetAllMessagesGetResponseable instead.
type ItemTeamChannelsGetallmessagesGetAllMessagesResponse struct {
    ItemTeamChannelsGetallmessagesGetAllMessagesGetResponse
}
// NewItemTeamChannelsGetallmessagesGetAllMessagesResponse instantiates a new ItemTeamChannelsGetallmessagesGetAllMessagesResponse and sets the default values.
func NewItemTeamChannelsGetallmessagesGetAllMessagesResponse()(*ItemTeamChannelsGetallmessagesGetAllMessagesResponse) {
    m := &ItemTeamChannelsGetallmessagesGetAllMessagesResponse{
        ItemTeamChannelsGetallmessagesGetAllMessagesGetResponse: *NewItemTeamChannelsGetallmessagesGetAllMessagesGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsGetallmessagesGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsGetallmessagesGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsGetallmessagesGetAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsGetallmessagesGetAllMessagesGetResponseable instead.
type ItemTeamChannelsGetallmessagesGetAllMessagesResponseable interface {
    ItemTeamChannelsGetallmessagesGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

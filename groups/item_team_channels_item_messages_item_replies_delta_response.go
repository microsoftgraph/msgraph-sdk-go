package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamChannelsItemMessagesItemRepliesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamChannelsItemMessagesItemRepliesDeltaResponse struct {
    ItemTeamChannelsItemMessagesItemRepliesDeltaGetResponse
}
// NewItemTeamChannelsItemMessagesItemRepliesDeltaResponse instantiates a new ItemTeamChannelsItemMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemTeamChannelsItemMessagesItemRepliesDeltaResponse()(*ItemTeamChannelsItemMessagesItemRepliesDeltaResponse) {
    m := &ItemTeamChannelsItemMessagesItemRepliesDeltaResponse{
        ItemTeamChannelsItemMessagesItemRepliesDeltaGetResponse: *NewItemTeamChannelsItemMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemMessagesItemRepliesDeltaResponse(), nil
}
// ItemTeamChannelsItemMessagesItemRepliesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamChannelsItemMessagesItemRepliesDeltaResponseable interface {
    ItemTeamChannelsItemMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

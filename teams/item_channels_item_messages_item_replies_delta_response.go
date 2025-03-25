package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChannelsItemMessagesItemRepliesDeltagetResponseable instead.
type ItemChannelsItemMessagesItemRepliesDeltaResponse struct {
    ItemChannelsItemMessagesItemRepliesDeltagetResponse
}
// NewItemChannelsItemMessagesItemRepliesDeltaResponse instantiates a new ItemChannelsItemMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemChannelsItemMessagesItemRepliesDeltaResponse()(*ItemChannelsItemMessagesItemRepliesDeltaResponse) {
    m := &ItemChannelsItemMessagesItemRepliesDeltaResponse{
        ItemChannelsItemMessagesItemRepliesDeltagetResponse: *NewItemChannelsItemMessagesItemRepliesDeltagetResponse(),
    }
    return m
}
// CreateItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemMessagesItemRepliesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChannelsItemMessagesItemRepliesDeltagetResponseable instead.
type ItemChannelsItemMessagesItemRepliesDeltaResponseable interface {
    ItemChannelsItemMessagesItemRepliesDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

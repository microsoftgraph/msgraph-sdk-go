package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChannelsItemMessagesItemRepliesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemChannelsItemMessagesItemRepliesDeltaResponse struct {
    ItemChannelsItemMessagesItemRepliesDeltaGetResponse
}
// NewItemChannelsItemMessagesItemRepliesDeltaResponse instantiates a new ItemChannelsItemMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemChannelsItemMessagesItemRepliesDeltaResponse()(*ItemChannelsItemMessagesItemRepliesDeltaResponse) {
    m := &ItemChannelsItemMessagesItemRepliesDeltaResponse{
        ItemChannelsItemMessagesItemRepliesDeltaGetResponse: *NewItemChannelsItemMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsItemMessagesItemRepliesDeltaResponse(), nil
}
// ItemChannelsItemMessagesItemRepliesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemChannelsItemMessagesItemRepliesDeltaResponseable interface {
    ItemChannelsItemMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

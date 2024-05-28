package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimarychannelMessagesItemRepliesDeltaGetResponseable instead.
type ItemPrimarychannelMessagesItemRepliesDeltaResponse struct {
    ItemPrimarychannelMessagesItemRepliesDeltaGetResponse
}
// NewItemPrimarychannelMessagesItemRepliesDeltaResponse instantiates a new ItemPrimarychannelMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemPrimarychannelMessagesItemRepliesDeltaResponse()(*ItemPrimarychannelMessagesItemRepliesDeltaResponse) {
    m := &ItemPrimarychannelMessagesItemRepliesDeltaResponse{
        ItemPrimarychannelMessagesItemRepliesDeltaGetResponse: *NewItemPrimarychannelMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateItemPrimarychannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimarychannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimarychannelMessagesItemRepliesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimarychannelMessagesItemRepliesDeltaGetResponseable instead.
type ItemPrimarychannelMessagesItemRepliesDeltaResponseable interface {
    ItemPrimarychannelMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

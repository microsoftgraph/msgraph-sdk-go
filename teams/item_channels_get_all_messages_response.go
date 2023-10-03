package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChannelsGetAllMessagesResponse 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemChannelsGetAllMessagesResponse struct {
    ItemChannelsGetAllMessagesGetResponse
}
// NewItemChannelsGetAllMessagesResponse instantiates a new ItemChannelsGetAllMessagesResponse and sets the default values.
func NewItemChannelsGetAllMessagesResponse()(*ItemChannelsGetAllMessagesResponse) {
    m := &ItemChannelsGetAllMessagesResponse{
        ItemChannelsGetAllMessagesGetResponse: *NewItemChannelsGetAllMessagesGetResponse(),
    }
    return m
}
// CreateItemChannelsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChannelsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChannelsGetAllMessagesResponse(), nil
}
// ItemChannelsGetAllMessagesResponseable 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemChannelsGetAllMessagesResponseable interface {
    ItemChannelsGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

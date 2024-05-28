package chats

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetallmessagesGetAllMessagesGetResponseable instead.
type GetallmessagesGetAllMessagesResponse struct {
    GetallmessagesGetAllMessagesGetResponse
}
// NewGetallmessagesGetAllMessagesResponse instantiates a new GetallmessagesGetAllMessagesResponse and sets the default values.
func NewGetallmessagesGetAllMessagesResponse()(*GetallmessagesGetAllMessagesResponse) {
    m := &GetallmessagesGetAllMessagesResponse{
        GetallmessagesGetAllMessagesGetResponse: *NewGetallmessagesGetAllMessagesGetResponse(),
    }
    return m
}
// CreateGetallmessagesGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetallmessagesGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetallmessagesGetAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use GetallmessagesGetAllMessagesGetResponseable instead.
type GetallmessagesGetAllMessagesResponseable interface {
    GetallmessagesGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

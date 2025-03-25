package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemChatsGetAllMessagesgetResponseable instead.
type ItemChatsGetAllMessagesResponse struct {
    ItemChatsGetAllMessagesgetResponse
}
// NewItemChatsGetAllMessagesResponse instantiates a new ItemChatsGetAllMessagesResponse and sets the default values.
func NewItemChatsGetAllMessagesResponse()(*ItemChatsGetAllMessagesResponse) {
    m := &ItemChatsGetAllMessagesResponse{
        ItemChatsGetAllMessagesgetResponse: *NewItemChatsGetAllMessagesgetResponse(),
    }
    return m
}
// CreateItemChatsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemChatsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsGetAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemChatsGetAllMessagesgetResponseable instead.
type ItemChatsGetAllMessagesResponseable interface {
    ItemChatsGetAllMessagesgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

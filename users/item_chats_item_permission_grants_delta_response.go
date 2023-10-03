package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemPermissionGrantsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemChatsItemPermissionGrantsDeltaResponse struct {
    ItemChatsItemPermissionGrantsDeltaGetResponse
}
// NewItemChatsItemPermissionGrantsDeltaResponse instantiates a new ItemChatsItemPermissionGrantsDeltaResponse and sets the default values.
func NewItemChatsItemPermissionGrantsDeltaResponse()(*ItemChatsItemPermissionGrantsDeltaResponse) {
    m := &ItemChatsItemPermissionGrantsDeltaResponse{
        ItemChatsItemPermissionGrantsDeltaGetResponse: *NewItemChatsItemPermissionGrantsDeltaGetResponse(),
    }
    return m
}
// CreateItemChatsItemPermissionGrantsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemPermissionGrantsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemPermissionGrantsDeltaResponse(), nil
}
// ItemChatsItemPermissionGrantsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemChatsItemPermissionGrantsDeltaResponseable interface {
    ItemChatsItemPermissionGrantsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

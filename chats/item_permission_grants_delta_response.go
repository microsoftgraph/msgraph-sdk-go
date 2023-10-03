package chats

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPermissionGrantsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPermissionGrantsDeltaResponse struct {
    ItemPermissionGrantsDeltaGetResponse
}
// NewItemPermissionGrantsDeltaResponse instantiates a new ItemPermissionGrantsDeltaResponse and sets the default values.
func NewItemPermissionGrantsDeltaResponse()(*ItemPermissionGrantsDeltaResponse) {
    m := &ItemPermissionGrantsDeltaResponse{
        ItemPermissionGrantsDeltaGetResponse: *NewItemPermissionGrantsDeltaGetResponse(),
    }
    return m
}
// CreateItemPermissionGrantsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsDeltaResponse(), nil
}
// ItemPermissionGrantsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPermissionGrantsDeltaResponseable interface {
    ItemPermissionGrantsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

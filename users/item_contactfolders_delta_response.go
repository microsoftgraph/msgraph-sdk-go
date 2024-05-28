package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContactfoldersDeltaGetResponseable instead.
type ItemContactfoldersDeltaResponse struct {
    ItemContactfoldersDeltaGetResponse
}
// NewItemContactfoldersDeltaResponse instantiates a new ItemContactfoldersDeltaResponse and sets the default values.
func NewItemContactfoldersDeltaResponse()(*ItemContactfoldersDeltaResponse) {
    m := &ItemContactfoldersDeltaResponse{
        ItemContactfoldersDeltaGetResponse: *NewItemContactfoldersDeltaGetResponse(),
    }
    return m
}
// CreateItemContactfoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContactfoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactfoldersDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContactfoldersDeltaGetResponseable instead.
type ItemContactfoldersDeltaResponseable interface {
    ItemContactfoldersDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

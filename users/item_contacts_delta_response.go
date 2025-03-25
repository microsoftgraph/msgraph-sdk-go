package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContactsDeltagetResponseable instead.
type ItemContactsDeltaResponse struct {
    ItemContactsDeltagetResponse
}
// NewItemContactsDeltaResponse instantiates a new ItemContactsDeltaResponse and sets the default values.
func NewItemContactsDeltaResponse()(*ItemContactsDeltaResponse) {
    m := &ItemContactsDeltaResponse{
        ItemContactsDeltagetResponse: *NewItemContactsDeltagetResponse(),
    }
    return m
}
// CreateItemContactsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContactsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContactsDeltagetResponseable instead.
type ItemContactsDeltaResponseable interface {
    ItemContactsDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContactfoldersItemContactsDeltaGetResponseable instead.
type ItemContactfoldersItemContactsDeltaResponse struct {
    ItemContactfoldersItemContactsDeltaGetResponse
}
// NewItemContactfoldersItemContactsDeltaResponse instantiates a new ItemContactfoldersItemContactsDeltaResponse and sets the default values.
func NewItemContactfoldersItemContactsDeltaResponse()(*ItemContactfoldersItemContactsDeltaResponse) {
    m := &ItemContactfoldersItemContactsDeltaResponse{
        ItemContactfoldersItemContactsDeltaGetResponse: *NewItemContactfoldersItemContactsDeltaGetResponse(),
    }
    return m
}
// CreateItemContactfoldersItemContactsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContactfoldersItemContactsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactfoldersItemContactsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContactfoldersItemContactsDeltaGetResponseable instead.
type ItemContactfoldersItemContactsDeltaResponseable interface {
    ItemContactfoldersItemContactsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContactfoldersItemChildfoldersDeltaGetResponseable instead.
type ItemContactfoldersItemChildfoldersDeltaResponse struct {
    ItemContactfoldersItemChildfoldersDeltaGetResponse
}
// NewItemContactfoldersItemChildfoldersDeltaResponse instantiates a new ItemContactfoldersItemChildfoldersDeltaResponse and sets the default values.
func NewItemContactfoldersItemChildfoldersDeltaResponse()(*ItemContactfoldersItemChildfoldersDeltaResponse) {
    m := &ItemContactfoldersItemChildfoldersDeltaResponse{
        ItemContactfoldersItemChildfoldersDeltaGetResponse: *NewItemContactfoldersItemChildfoldersDeltaGetResponse(),
    }
    return m
}
// CreateItemContactfoldersItemChildfoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContactfoldersItemChildfoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactfoldersItemChildfoldersDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContactfoldersItemChildfoldersDeltaGetResponseable instead.
type ItemContactfoldersItemChildfoldersDeltaResponseable interface {
    ItemContactfoldersItemChildfoldersDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

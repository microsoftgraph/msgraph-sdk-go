package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContactFoldersItemChildFoldersItemContactsDeltagetResponseable instead.
type ItemContactFoldersItemChildFoldersItemContactsDeltaResponse struct {
    ItemContactFoldersItemChildFoldersItemContactsDeltagetResponse
}
// NewItemContactFoldersItemChildFoldersItemContactsDeltaResponse instantiates a new ItemContactFoldersItemChildFoldersItemContactsDeltaResponse and sets the default values.
func NewItemContactFoldersItemChildFoldersItemContactsDeltaResponse()(*ItemContactFoldersItemChildFoldersItemContactsDeltaResponse) {
    m := &ItemContactFoldersItemChildFoldersItemContactsDeltaResponse{
        ItemContactFoldersItemChildFoldersItemContactsDeltagetResponse: *NewItemContactFoldersItemChildFoldersItemContactsDeltagetResponse(),
    }
    return m
}
// CreateItemContactFoldersItemChildFoldersItemContactsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContactFoldersItemChildFoldersItemContactsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactFoldersItemChildFoldersItemContactsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContactFoldersItemChildFoldersItemContactsDeltagetResponseable instead.
type ItemContactFoldersItemChildFoldersItemContactsDeltaResponseable interface {
    ItemContactFoldersItemChildFoldersItemContactsDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

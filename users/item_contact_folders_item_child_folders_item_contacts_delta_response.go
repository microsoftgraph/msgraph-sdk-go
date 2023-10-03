package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemContactFoldersItemChildFoldersItemContactsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemContactFoldersItemChildFoldersItemContactsDeltaResponse struct {
    ItemContactFoldersItemChildFoldersItemContactsDeltaGetResponse
}
// NewItemContactFoldersItemChildFoldersItemContactsDeltaResponse instantiates a new ItemContactFoldersItemChildFoldersItemContactsDeltaResponse and sets the default values.
func NewItemContactFoldersItemChildFoldersItemContactsDeltaResponse()(*ItemContactFoldersItemChildFoldersItemContactsDeltaResponse) {
    m := &ItemContactFoldersItemChildFoldersItemContactsDeltaResponse{
        ItemContactFoldersItemChildFoldersItemContactsDeltaGetResponse: *NewItemContactFoldersItemChildFoldersItemContactsDeltaGetResponse(),
    }
    return m
}
// CreateItemContactFoldersItemChildFoldersItemContactsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemContactFoldersItemChildFoldersItemContactsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactFoldersItemChildFoldersItemContactsDeltaResponse(), nil
}
// ItemContactFoldersItemChildFoldersItemContactsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemContactFoldersItemChildFoldersItemContactsDeltaResponseable interface {
    ItemContactFoldersItemChildFoldersItemContactsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

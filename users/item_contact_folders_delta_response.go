package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemContactFoldersDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemContactFoldersDeltaResponse struct {
    ItemContactFoldersDeltaGetResponse
}
// NewItemContactFoldersDeltaResponse instantiates a new ItemContactFoldersDeltaResponse and sets the default values.
func NewItemContactFoldersDeltaResponse()(*ItemContactFoldersDeltaResponse) {
    m := &ItemContactFoldersDeltaResponse{
        ItemContactFoldersDeltaGetResponse: *NewItemContactFoldersDeltaGetResponse(),
    }
    return m
}
// CreateItemContactFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemContactFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactFoldersDeltaResponse(), nil
}
// ItemContactFoldersDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemContactFoldersDeltaResponseable interface {
    ItemContactFoldersDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMailFoldersDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMailFoldersDeltaResponse struct {
    ItemMailFoldersDeltaGetResponse
}
// NewItemMailFoldersDeltaResponse instantiates a new ItemMailFoldersDeltaResponse and sets the default values.
func NewItemMailFoldersDeltaResponse()(*ItemMailFoldersDeltaResponse) {
    m := &ItemMailFoldersDeltaResponse{
        ItemMailFoldersDeltaGetResponse: *NewItemMailFoldersDeltaGetResponse(),
    }
    return m
}
// CreateItemMailFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMailFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailFoldersDeltaResponse(), nil
}
// ItemMailFoldersDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMailFoldersDeltaResponseable interface {
    ItemMailFoldersDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

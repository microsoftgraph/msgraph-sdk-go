package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMailFoldersItemMessagesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMailFoldersItemMessagesDeltaResponse struct {
    ItemMailFoldersItemMessagesDeltaGetResponse
}
// NewItemMailFoldersItemMessagesDeltaResponse instantiates a new ItemMailFoldersItemMessagesDeltaResponse and sets the default values.
func NewItemMailFoldersItemMessagesDeltaResponse()(*ItemMailFoldersItemMessagesDeltaResponse) {
    m := &ItemMailFoldersItemMessagesDeltaResponse{
        ItemMailFoldersItemMessagesDeltaGetResponse: *NewItemMailFoldersItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemMailFoldersItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMailFoldersItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailFoldersItemMessagesDeltaResponse(), nil
}
// ItemMailFoldersItemMessagesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMailFoldersItemMessagesDeltaResponseable interface {
    ItemMailFoldersItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMailFoldersItemChildFoldersDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMailFoldersItemChildFoldersDeltaResponse struct {
    ItemMailFoldersItemChildFoldersDeltaGetResponse
}
// NewItemMailFoldersItemChildFoldersDeltaResponse instantiates a new ItemMailFoldersItemChildFoldersDeltaResponse and sets the default values.
func NewItemMailFoldersItemChildFoldersDeltaResponse()(*ItemMailFoldersItemChildFoldersDeltaResponse) {
    m := &ItemMailFoldersItemChildFoldersDeltaResponse{
        ItemMailFoldersItemChildFoldersDeltaGetResponse: *NewItemMailFoldersItemChildFoldersDeltaGetResponse(),
    }
    return m
}
// CreateItemMailFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMailFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailFoldersItemChildFoldersDeltaResponse(), nil
}
// ItemMailFoldersItemChildFoldersDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMailFoldersItemChildFoldersDeltaResponseable interface {
    ItemMailFoldersItemChildFoldersDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

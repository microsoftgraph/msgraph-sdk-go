package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemMailFoldersItemChildFoldersDeltagetResponseable instead.
type ItemMailFoldersItemChildFoldersDeltaResponse struct {
    ItemMailFoldersItemChildFoldersDeltagetResponse
}
// NewItemMailFoldersItemChildFoldersDeltaResponse instantiates a new ItemMailFoldersItemChildFoldersDeltaResponse and sets the default values.
func NewItemMailFoldersItemChildFoldersDeltaResponse()(*ItemMailFoldersItemChildFoldersDeltaResponse) {
    m := &ItemMailFoldersItemChildFoldersDeltaResponse{
        ItemMailFoldersItemChildFoldersDeltagetResponse: *NewItemMailFoldersItemChildFoldersDeltagetResponse(),
    }
    return m
}
// CreateItemMailFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemMailFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailFoldersItemChildFoldersDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemMailFoldersItemChildFoldersDeltagetResponseable instead.
type ItemMailFoldersItemChildFoldersDeltaResponseable interface {
    ItemMailFoldersItemChildFoldersDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

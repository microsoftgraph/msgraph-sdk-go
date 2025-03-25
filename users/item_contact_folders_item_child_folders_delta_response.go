package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContactFoldersItemChildFoldersDeltagetResponseable instead.
type ItemContactFoldersItemChildFoldersDeltaResponse struct {
    ItemContactFoldersItemChildFoldersDeltagetResponse
}
// NewItemContactFoldersItemChildFoldersDeltaResponse instantiates a new ItemContactFoldersItemChildFoldersDeltaResponse and sets the default values.
func NewItemContactFoldersItemChildFoldersDeltaResponse()(*ItemContactFoldersItemChildFoldersDeltaResponse) {
    m := &ItemContactFoldersItemChildFoldersDeltaResponse{
        ItemContactFoldersItemChildFoldersDeltagetResponse: *NewItemContactFoldersItemChildFoldersDeltagetResponse(),
    }
    return m
}
// CreateItemContactFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContactFoldersItemChildFoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactFoldersItemChildFoldersDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContactFoldersItemChildFoldersDeltagetResponseable instead.
type ItemContactFoldersItemChildFoldersDeltaResponseable interface {
    ItemContactFoldersItemChildFoldersDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

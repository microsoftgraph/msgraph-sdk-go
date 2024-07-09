package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemInvitePostResponseable instead.
type FileStorageContainersItemDriveItemsItemInviteResponse struct {
    FileStorageContainersItemDriveItemsItemInvitePostResponse
}
// NewFileStorageContainersItemDriveItemsItemInviteResponse instantiates a new FileStorageContainersItemDriveItemsItemInviteResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemInviteResponse()(*FileStorageContainersItemDriveItemsItemInviteResponse) {
    m := &FileStorageContainersItemDriveItemsItemInviteResponse{
        FileStorageContainersItemDriveItemsItemInvitePostResponse: *NewFileStorageContainersItemDriveItemsItemInvitePostResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemInviteResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemInvitePostResponseable instead.
type FileStorageContainersItemDriveItemsItemInviteResponseable interface {
    FileStorageContainersItemDriveItemsItemInvitePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

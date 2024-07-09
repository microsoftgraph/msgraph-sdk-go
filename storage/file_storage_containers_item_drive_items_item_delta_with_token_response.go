package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemDeltaWithTokenGetResponseable instead.
type FileStorageContainersItemDriveItemsItemDeltaWithTokenResponse struct {
    FileStorageContainersItemDriveItemsItemDeltaWithTokenGetResponse
}
// NewFileStorageContainersItemDriveItemsItemDeltaWithTokenResponse instantiates a new FileStorageContainersItemDriveItemsItemDeltaWithTokenResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemDeltaWithTokenResponse()(*FileStorageContainersItemDriveItemsItemDeltaWithTokenResponse) {
    m := &FileStorageContainersItemDriveItemsItemDeltaWithTokenResponse{
        FileStorageContainersItemDriveItemsItemDeltaWithTokenGetResponse: *NewFileStorageContainersItemDriveItemsItemDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemDeltaWithTokenGetResponseable instead.
type FileStorageContainersItemDriveItemsItemDeltaWithTokenResponseable interface {
    FileStorageContainersItemDriveItemsItemDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

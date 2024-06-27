package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemSearchWithQGetResponseable instead.
type FileStorageContainersItemDriveItemsItemSearchWithQResponse struct {
    FileStorageContainersItemDriveItemsItemSearchWithQGetResponse
}
// NewFileStorageContainersItemDriveItemsItemSearchWithQResponse instantiates a new FileStorageContainersItemDriveItemsItemSearchWithQResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemSearchWithQResponse()(*FileStorageContainersItemDriveItemsItemSearchWithQResponse) {
    m := &FileStorageContainersItemDriveItemsItemSearchWithQResponse{
        FileStorageContainersItemDriveItemsItemSearchWithQGetResponse: *NewFileStorageContainersItemDriveItemsItemSearchWithQGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemSearchWithQResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemSearchWithQResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemSearchWithQResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemSearchWithQGetResponseable instead.
type FileStorageContainersItemDriveItemsItemSearchWithQResponseable interface {
    FileStorageContainersItemDriveItemsItemSearchWithQGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

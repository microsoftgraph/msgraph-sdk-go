package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveSearchWithQGetResponseable instead.
type FileStorageContainersItemDriveSearchWithQResponse struct {
    FileStorageContainersItemDriveSearchWithQGetResponse
}
// NewFileStorageContainersItemDriveSearchWithQResponse instantiates a new FileStorageContainersItemDriveSearchWithQResponse and sets the default values.
func NewFileStorageContainersItemDriveSearchWithQResponse()(*FileStorageContainersItemDriveSearchWithQResponse) {
    m := &FileStorageContainersItemDriveSearchWithQResponse{
        FileStorageContainersItemDriveSearchWithQGetResponse: *NewFileStorageContainersItemDriveSearchWithQGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveSearchWithQResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveSearchWithQResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveSearchWithQResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveSearchWithQGetResponseable instead.
type FileStorageContainersItemDriveSearchWithQResponseable interface {
    FileStorageContainersItemDriveSearchWithQGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

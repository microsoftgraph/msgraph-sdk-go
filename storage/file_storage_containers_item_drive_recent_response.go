package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveRecentGetResponseable instead.
type FileStorageContainersItemDriveRecentResponse struct {
    FileStorageContainersItemDriveRecentGetResponse
}
// NewFileStorageContainersItemDriveRecentResponse instantiates a new FileStorageContainersItemDriveRecentResponse and sets the default values.
func NewFileStorageContainersItemDriveRecentResponse()(*FileStorageContainersItemDriveRecentResponse) {
    m := &FileStorageContainersItemDriveRecentResponse{
        FileStorageContainersItemDriveRecentGetResponse: *NewFileStorageContainersItemDriveRecentGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveRecentResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveRecentResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveRecentResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveRecentGetResponseable instead.
type FileStorageContainersItemDriveRecentResponseable interface {
    FileStorageContainersItemDriveRecentGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

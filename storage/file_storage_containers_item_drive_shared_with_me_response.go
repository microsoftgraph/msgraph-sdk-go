package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveSharedWithMeGetResponseable instead.
type FileStorageContainersItemDriveSharedWithMeResponse struct {
    FileStorageContainersItemDriveSharedWithMeGetResponse
}
// NewFileStorageContainersItemDriveSharedWithMeResponse instantiates a new FileStorageContainersItemDriveSharedWithMeResponse and sets the default values.
func NewFileStorageContainersItemDriveSharedWithMeResponse()(*FileStorageContainersItemDriveSharedWithMeResponse) {
    m := &FileStorageContainersItemDriveSharedWithMeResponse{
        FileStorageContainersItemDriveSharedWithMeGetResponse: *NewFileStorageContainersItemDriveSharedWithMeGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveSharedWithMeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveSharedWithMeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveSharedWithMeResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveSharedWithMeGetResponseable instead.
type FileStorageContainersItemDriveSharedWithMeResponseable interface {
    FileStorageContainersItemDriveSharedWithMeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

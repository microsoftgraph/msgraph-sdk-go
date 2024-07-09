package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListItemsDeltaGetResponseable instead.
type FileStorageContainersItemDriveListItemsDeltaResponse struct {
    FileStorageContainersItemDriveListItemsDeltaGetResponse
}
// NewFileStorageContainersItemDriveListItemsDeltaResponse instantiates a new FileStorageContainersItemDriveListItemsDeltaResponse and sets the default values.
func NewFileStorageContainersItemDriveListItemsDeltaResponse()(*FileStorageContainersItemDriveListItemsDeltaResponse) {
    m := &FileStorageContainersItemDriveListItemsDeltaResponse{
        FileStorageContainersItemDriveListItemsDeltaGetResponse: *NewFileStorageContainersItemDriveListItemsDeltaGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveListItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveListItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveListItemsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListItemsDeltaGetResponseable instead.
type FileStorageContainersItemDriveListItemsDeltaResponseable interface {
    FileStorageContainersItemDriveListItemsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

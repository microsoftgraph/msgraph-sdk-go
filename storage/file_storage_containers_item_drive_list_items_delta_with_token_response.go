package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListItemsDeltaWithTokenGetResponseable instead.
type FileStorageContainersItemDriveListItemsDeltaWithTokenResponse struct {
    FileStorageContainersItemDriveListItemsDeltaWithTokenGetResponse
}
// NewFileStorageContainersItemDriveListItemsDeltaWithTokenResponse instantiates a new FileStorageContainersItemDriveListItemsDeltaWithTokenResponse and sets the default values.
func NewFileStorageContainersItemDriveListItemsDeltaWithTokenResponse()(*FileStorageContainersItemDriveListItemsDeltaWithTokenResponse) {
    m := &FileStorageContainersItemDriveListItemsDeltaWithTokenResponse{
        FileStorageContainersItemDriveListItemsDeltaWithTokenGetResponse: *NewFileStorageContainersItemDriveListItemsDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveListItemsDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveListItemsDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveListItemsDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListItemsDeltaWithTokenGetResponseable instead.
type FileStorageContainersItemDriveListItemsDeltaWithTokenResponseable interface {
    FileStorageContainersItemDriveListItemsDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

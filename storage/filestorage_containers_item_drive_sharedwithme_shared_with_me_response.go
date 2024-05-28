package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponseable instead.
type FilestorageContainersItemDriveSharedwithmeSharedWithMeResponse struct {
    FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponse
}
// NewFilestorageContainersItemDriveSharedwithmeSharedWithMeResponse instantiates a new FilestorageContainersItemDriveSharedwithmeSharedWithMeResponse and sets the default values.
func NewFilestorageContainersItemDriveSharedwithmeSharedWithMeResponse()(*FilestorageContainersItemDriveSharedwithmeSharedWithMeResponse) {
    m := &FilestorageContainersItemDriveSharedwithmeSharedWithMeResponse{
        FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponse: *NewFilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveSharedwithmeSharedWithMeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveSharedwithmeSharedWithMeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveSharedwithmeSharedWithMeResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponseable instead.
type FilestorageContainersItemDriveSharedwithmeSharedWithMeResponseable interface {
    FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

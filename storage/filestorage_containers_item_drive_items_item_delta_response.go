package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemDeltaGetResponseable instead.
type FilestorageContainersItemDriveItemsItemDeltaResponse struct {
    FilestorageContainersItemDriveItemsItemDeltaGetResponse
}
// NewFilestorageContainersItemDriveItemsItemDeltaResponse instantiates a new FilestorageContainersItemDriveItemsItemDeltaResponse and sets the default values.
func NewFilestorageContainersItemDriveItemsItemDeltaResponse()(*FilestorageContainersItemDriveItemsItemDeltaResponse) {
    m := &FilestorageContainersItemDriveItemsItemDeltaResponse{
        FilestorageContainersItemDriveItemsItemDeltaGetResponse: *NewFilestorageContainersItemDriveItemsItemDeltaGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveItemsItemDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveItemsItemDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveItemsItemDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemDeltaGetResponseable instead.
type FilestorageContainersItemDriveItemsItemDeltaResponseable interface {
    FilestorageContainersItemDriveItemsItemDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

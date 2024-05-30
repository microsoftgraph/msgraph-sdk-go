package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListItemsDeltaGetResponseable instead.
type FilestorageContainersItemDriveListItemsDeltaResponse struct {
    FilestorageContainersItemDriveListItemsDeltaGetResponse
}
// NewFilestorageContainersItemDriveListItemsDeltaResponse instantiates a new FilestorageContainersItemDriveListItemsDeltaResponse and sets the default values.
func NewFilestorageContainersItemDriveListItemsDeltaResponse()(*FilestorageContainersItemDriveListItemsDeltaResponse) {
    m := &FilestorageContainersItemDriveListItemsDeltaResponse{
        FilestorageContainersItemDriveListItemsDeltaGetResponse: *NewFilestorageContainersItemDriveListItemsDeltaGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveListItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveListItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveListItemsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListItemsDeltaGetResponseable instead.
type FilestorageContainersItemDriveListItemsDeltaResponseable interface {
    FilestorageContainersItemDriveListItemsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

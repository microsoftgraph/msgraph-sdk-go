package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveRecentGetResponseable instead.
type FilestorageContainersItemDriveRecentResponse struct {
    FilestorageContainersItemDriveRecentGetResponse
}
// NewFilestorageContainersItemDriveRecentResponse instantiates a new FilestorageContainersItemDriveRecentResponse and sets the default values.
func NewFilestorageContainersItemDriveRecentResponse()(*FilestorageContainersItemDriveRecentResponse) {
    m := &FilestorageContainersItemDriveRecentResponse{
        FilestorageContainersItemDriveRecentGetResponse: *NewFilestorageContainersItemDriveRecentGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveRecentResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveRecentResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveRecentResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveRecentGetResponseable instead.
type FilestorageContainersItemDriveRecentResponseable interface {
    FilestorageContainersItemDriveRecentGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

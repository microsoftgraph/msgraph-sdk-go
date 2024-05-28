package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveSearchwithqSearchWithQGetResponseable instead.
type FilestorageContainersItemDriveSearchwithqSearchWithQResponse struct {
    FilestorageContainersItemDriveSearchwithqSearchWithQGetResponse
}
// NewFilestorageContainersItemDriveSearchwithqSearchWithQResponse instantiates a new FilestorageContainersItemDriveSearchwithqSearchWithQResponse and sets the default values.
func NewFilestorageContainersItemDriveSearchwithqSearchWithQResponse()(*FilestorageContainersItemDriveSearchwithqSearchWithQResponse) {
    m := &FilestorageContainersItemDriveSearchwithqSearchWithQResponse{
        FilestorageContainersItemDriveSearchwithqSearchWithQGetResponse: *NewFilestorageContainersItemDriveSearchwithqSearchWithQGetResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveSearchwithqSearchWithQResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveSearchwithqSearchWithQResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveSearchwithqSearchWithQResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveSearchwithqSearchWithQGetResponseable instead.
type FilestorageContainersItemDriveSearchwithqSearchWithQResponseable interface {
    FilestorageContainersItemDriveSearchwithqSearchWithQGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemInvitePostResponseable instead.
type FilestorageContainersItemDriveItemsItemInviteResponse struct {
    FilestorageContainersItemDriveItemsItemInvitePostResponse
}
// NewFilestorageContainersItemDriveItemsItemInviteResponse instantiates a new FilestorageContainersItemDriveItemsItemInviteResponse and sets the default values.
func NewFilestorageContainersItemDriveItemsItemInviteResponse()(*FilestorageContainersItemDriveItemsItemInviteResponse) {
    m := &FilestorageContainersItemDriveItemsItemInviteResponse{
        FilestorageContainersItemDriveItemsItemInvitePostResponse: *NewFilestorageContainersItemDriveItemsItemInvitePostResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveItemsItemInviteResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveItemsItemInvitePostResponseable instead.
type FilestorageContainersItemDriveItemsItemInviteResponseable interface {
    FilestorageContainersItemDriveItemsItemInvitePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemPermissionsItemGrantPostResponseable instead.
type FilestorageContainersItemPermissionsItemGrantResponse struct {
    FilestorageContainersItemPermissionsItemGrantPostResponse
}
// NewFilestorageContainersItemPermissionsItemGrantResponse instantiates a new FilestorageContainersItemPermissionsItemGrantResponse and sets the default values.
func NewFilestorageContainersItemPermissionsItemGrantResponse()(*FilestorageContainersItemPermissionsItemGrantResponse) {
    m := &FilestorageContainersItemPermissionsItemGrantResponse{
        FilestorageContainersItemPermissionsItemGrantPostResponse: *NewFilestorageContainersItemPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateFilestorageContainersItemPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemPermissionsItemGrantPostResponseable instead.
type FilestorageContainersItemPermissionsItemGrantResponseable interface {
    FilestorageContainersItemPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

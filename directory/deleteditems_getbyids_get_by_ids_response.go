package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeleteditemsGetbyidsGetByIdsPostResponseable instead.
type DeleteditemsGetbyidsGetByIdsResponse struct {
    DeleteditemsGetbyidsGetByIdsPostResponse
}
// NewDeleteditemsGetbyidsGetByIdsResponse instantiates a new DeleteditemsGetbyidsGetByIdsResponse and sets the default values.
func NewDeleteditemsGetbyidsGetByIdsResponse()(*DeleteditemsGetbyidsGetByIdsResponse) {
    m := &DeleteditemsGetbyidsGetByIdsResponse{
        DeleteditemsGetbyidsGetByIdsPostResponse: *NewDeleteditemsGetbyidsGetByIdsPostResponse(),
    }
    return m
}
// CreateDeleteditemsGetbyidsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeleteditemsGetbyidsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleteditemsGetbyidsGetByIdsResponse(), nil
}
// Deprecated: This class is obsolete. Use DeleteditemsGetbyidsGetByIdsPostResponseable instead.
type DeleteditemsGetbyidsGetByIdsResponseable interface {
    DeleteditemsGetbyidsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

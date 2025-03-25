package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedItemsItemCheckMemberObjectspostResponseable instead.
type DeletedItemsItemCheckMemberObjectsResponse struct {
    DeletedItemsItemCheckMemberObjectspostResponse
}
// NewDeletedItemsItemCheckMemberObjectsResponse instantiates a new DeletedItemsItemCheckMemberObjectsResponse and sets the default values.
func NewDeletedItemsItemCheckMemberObjectsResponse()(*DeletedItemsItemCheckMemberObjectsResponse) {
    m := &DeletedItemsItemCheckMemberObjectsResponse{
        DeletedItemsItemCheckMemberObjectspostResponse: *NewDeletedItemsItemCheckMemberObjectspostResponse(),
    }
    return m
}
// CreateDeletedItemsItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedItemsItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedItemsItemCheckMemberObjectsResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedItemsItemCheckMemberObjectspostResponseable instead.
type DeletedItemsItemCheckMemberObjectsResponseable interface {
    DeletedItemsItemCheckMemberObjectspostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

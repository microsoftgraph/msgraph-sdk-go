package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletedItemsItemCheckMemberObjectsResponse 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type DeletedItemsItemCheckMemberObjectsResponse struct {
    DeletedItemsItemCheckMemberObjectsPostResponse
}
// NewDeletedItemsItemCheckMemberObjectsResponse instantiates a new DeletedItemsItemCheckMemberObjectsResponse and sets the default values.
func NewDeletedItemsItemCheckMemberObjectsResponse()(*DeletedItemsItemCheckMemberObjectsResponse) {
    m := &DeletedItemsItemCheckMemberObjectsResponse{
        DeletedItemsItemCheckMemberObjectsPostResponse: *NewDeletedItemsItemCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateDeletedItemsItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedItemsItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedItemsItemCheckMemberObjectsResponse(), nil
}
// DeletedItemsItemCheckMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type DeletedItemsItemCheckMemberObjectsResponseable interface {
    DeletedItemsItemCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletedItemsItemGetMemberObjectsResponse 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type DeletedItemsItemGetMemberObjectsResponse struct {
    DeletedItemsItemGetMemberObjectsPostResponse
}
// NewDeletedItemsItemGetMemberObjectsResponse instantiates a new DeletedItemsItemGetMemberObjectsResponse and sets the default values.
func NewDeletedItemsItemGetMemberObjectsResponse()(*DeletedItemsItemGetMemberObjectsResponse) {
    m := &DeletedItemsItemGetMemberObjectsResponse{
        DeletedItemsItemGetMemberObjectsPostResponse: *NewDeletedItemsItemGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateDeletedItemsItemGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedItemsItemGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedItemsItemGetMemberObjectsResponse(), nil
}
// DeletedItemsItemGetMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type DeletedItemsItemGetMemberObjectsResponseable interface {
    DeletedItemsItemGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

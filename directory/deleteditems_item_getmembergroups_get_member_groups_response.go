package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeleteditemsItemGetmembergroupsGetMemberGroupsPostResponseable instead.
type DeleteditemsItemGetmembergroupsGetMemberGroupsResponse struct {
    DeleteditemsItemGetmembergroupsGetMemberGroupsPostResponse
}
// NewDeleteditemsItemGetmembergroupsGetMemberGroupsResponse instantiates a new DeleteditemsItemGetmembergroupsGetMemberGroupsResponse and sets the default values.
func NewDeleteditemsItemGetmembergroupsGetMemberGroupsResponse()(*DeleteditemsItemGetmembergroupsGetMemberGroupsResponse) {
    m := &DeleteditemsItemGetmembergroupsGetMemberGroupsResponse{
        DeleteditemsItemGetmembergroupsGetMemberGroupsPostResponse: *NewDeleteditemsItemGetmembergroupsGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateDeleteditemsItemGetmembergroupsGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeleteditemsItemGetmembergroupsGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleteditemsItemGetmembergroupsGetMemberGroupsResponse(), nil
}
// Deprecated: This class is obsolete. Use DeleteditemsItemGetmembergroupsGetMemberGroupsPostResponseable instead.
type DeleteditemsItemGetmembergroupsGetMemberGroupsResponseable interface {
    DeleteditemsItemGetmembergroupsGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

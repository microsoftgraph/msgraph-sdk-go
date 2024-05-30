package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGetmembergroupsGetMemberGroupsPostResponseable instead.
type ItemGetmembergroupsGetMemberGroupsResponse struct {
    ItemGetmembergroupsGetMemberGroupsPostResponse
}
// NewItemGetmembergroupsGetMemberGroupsResponse instantiates a new ItemGetmembergroupsGetMemberGroupsResponse and sets the default values.
func NewItemGetmembergroupsGetMemberGroupsResponse()(*ItemGetmembergroupsGetMemberGroupsResponse) {
    m := &ItemGetmembergroupsGetMemberGroupsResponse{
        ItemGetmembergroupsGetMemberGroupsPostResponse: *NewItemGetmembergroupsGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemGetmembergroupsGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetmembergroupsGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetmembergroupsGetMemberGroupsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGetmembergroupsGetMemberGroupsPostResponseable instead.
type ItemGetmembergroupsGetMemberGroupsResponseable interface {
    ItemGetmembergroupsGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPermissionGrantsItemGetMemberGroupsResponse 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemPermissionGrantsItemGetMemberGroupsResponse struct {
    ItemPermissionGrantsItemGetMemberGroupsPostResponse
}
// NewItemPermissionGrantsItemGetMemberGroupsResponse instantiates a new ItemPermissionGrantsItemGetMemberGroupsResponse and sets the default values.
func NewItemPermissionGrantsItemGetMemberGroupsResponse()(*ItemPermissionGrantsItemGetMemberGroupsResponse) {
    m := &ItemPermissionGrantsItemGetMemberGroupsResponse{
        ItemPermissionGrantsItemGetMemberGroupsPostResponse: *NewItemPermissionGrantsItemGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsItemGetMemberGroupsResponse(), nil
}
// ItemPermissionGrantsItemGetMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemPermissionGrantsItemGetMemberGroupsResponseable interface {
    ItemPermissionGrantsItemGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

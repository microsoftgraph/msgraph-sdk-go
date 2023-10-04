package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemPermissionGrantsItemGetMemberGroupsResponse 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemChatsItemPermissionGrantsItemGetMemberGroupsResponse struct {
    ItemChatsItemPermissionGrantsItemGetMemberGroupsPostResponse
}
// NewItemChatsItemPermissionGrantsItemGetMemberGroupsResponse instantiates a new ItemChatsItemPermissionGrantsItemGetMemberGroupsResponse and sets the default values.
func NewItemChatsItemPermissionGrantsItemGetMemberGroupsResponse()(*ItemChatsItemPermissionGrantsItemGetMemberGroupsResponse) {
    m := &ItemChatsItemPermissionGrantsItemGetMemberGroupsResponse{
        ItemChatsItemPermissionGrantsItemGetMemberGroupsPostResponse: *NewItemChatsItemPermissionGrantsItemGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemChatsItemPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemPermissionGrantsItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemPermissionGrantsItemGetMemberGroupsResponse(), nil
}
// ItemChatsItemPermissionGrantsItemGetMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemChatsItemPermissionGrantsItemGetMemberGroupsResponseable interface {
    ItemChatsItemPermissionGrantsItemGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

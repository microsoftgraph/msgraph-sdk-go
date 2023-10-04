package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse struct {
    ItemChatsItemPermissionGrantsItemCheckMemberGroupsPostResponse
}
// NewItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse instantiates a new ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse and sets the default values.
func NewItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse()(*ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse) {
    m := &ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse{
        ItemChatsItemPermissionGrantsItemCheckMemberGroupsPostResponse: *NewItemChatsItemPermissionGrantsItemCheckMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemChatsItemPermissionGrantsItemCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemPermissionGrantsItemCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemPermissionGrantsItemCheckMemberGroupsResponse(), nil
}
// ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemChatsItemPermissionGrantsItemCheckMemberGroupsResponseable interface {
    ItemChatsItemPermissionGrantsItemCheckMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

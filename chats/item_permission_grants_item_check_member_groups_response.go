package chats

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPermissionGrantsItemCheckMemberGroupsResponse 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemPermissionGrantsItemCheckMemberGroupsResponse struct {
    ItemPermissionGrantsItemCheckMemberGroupsPostResponse
}
// NewItemPermissionGrantsItemCheckMemberGroupsResponse instantiates a new ItemPermissionGrantsItemCheckMemberGroupsResponse and sets the default values.
func NewItemPermissionGrantsItemCheckMemberGroupsResponse()(*ItemPermissionGrantsItemCheckMemberGroupsResponse) {
    m := &ItemPermissionGrantsItemCheckMemberGroupsResponse{
        ItemPermissionGrantsItemCheckMemberGroupsPostResponse: *NewItemPermissionGrantsItemCheckMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemPermissionGrantsItemCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsItemCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsItemCheckMemberGroupsResponse(), nil
}
// ItemPermissionGrantsItemCheckMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use checkMemberGroupsPostResponse instead.
type ItemPermissionGrantsItemCheckMemberGroupsResponseable interface {
    ItemPermissionGrantsItemCheckMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

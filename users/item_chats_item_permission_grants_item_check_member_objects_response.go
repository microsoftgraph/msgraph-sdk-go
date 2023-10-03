package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse struct {
    ItemChatsItemPermissionGrantsItemCheckMemberObjectsPostResponse
}
// NewItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse instantiates a new ItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse and sets the default values.
func NewItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse()(*ItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse) {
    m := &ItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse{
        ItemChatsItemPermissionGrantsItemCheckMemberObjectsPostResponse: *NewItemChatsItemPermissionGrantsItemCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemChatsItemPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemPermissionGrantsItemCheckMemberObjectsResponse(), nil
}
// ItemChatsItemPermissionGrantsItemCheckMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemChatsItemPermissionGrantsItemCheckMemberObjectsResponseable interface {
    ItemChatsItemPermissionGrantsItemCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

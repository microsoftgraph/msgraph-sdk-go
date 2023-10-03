package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemPermissionGrantsItemGetMemberObjectsResponse 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemChatsItemPermissionGrantsItemGetMemberObjectsResponse struct {
    ItemChatsItemPermissionGrantsItemGetMemberObjectsPostResponse
}
// NewItemChatsItemPermissionGrantsItemGetMemberObjectsResponse instantiates a new ItemChatsItemPermissionGrantsItemGetMemberObjectsResponse and sets the default values.
func NewItemChatsItemPermissionGrantsItemGetMemberObjectsResponse()(*ItemChatsItemPermissionGrantsItemGetMemberObjectsResponse) {
    m := &ItemChatsItemPermissionGrantsItemGetMemberObjectsResponse{
        ItemChatsItemPermissionGrantsItemGetMemberObjectsPostResponse: *NewItemChatsItemPermissionGrantsItemGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemChatsItemPermissionGrantsItemGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemPermissionGrantsItemGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemPermissionGrantsItemGetMemberObjectsResponse(), nil
}
// ItemChatsItemPermissionGrantsItemGetMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemChatsItemPermissionGrantsItemGetMemberObjectsResponseable interface {
    ItemChatsItemPermissionGrantsItemGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

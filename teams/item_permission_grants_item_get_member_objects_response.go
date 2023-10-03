package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPermissionGrantsItemGetMemberObjectsResponse 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemPermissionGrantsItemGetMemberObjectsResponse struct {
    ItemPermissionGrantsItemGetMemberObjectsPostResponse
}
// NewItemPermissionGrantsItemGetMemberObjectsResponse instantiates a new ItemPermissionGrantsItemGetMemberObjectsResponse and sets the default values.
func NewItemPermissionGrantsItemGetMemberObjectsResponse()(*ItemPermissionGrantsItemGetMemberObjectsResponse) {
    m := &ItemPermissionGrantsItemGetMemberObjectsResponse{
        ItemPermissionGrantsItemGetMemberObjectsPostResponse: *NewItemPermissionGrantsItemGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemPermissionGrantsItemGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsItemGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsItemGetMemberObjectsResponse(), nil
}
// ItemPermissionGrantsItemGetMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemPermissionGrantsItemGetMemberObjectsResponseable interface {
    ItemPermissionGrantsItemGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}

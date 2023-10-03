package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPermissionGrantsItemCheckMemberObjectsResponse 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemPermissionGrantsItemCheckMemberObjectsResponse struct {
    ItemPermissionGrantsItemCheckMemberObjectsPostResponse
}
// NewItemPermissionGrantsItemCheckMemberObjectsResponse instantiates a new ItemPermissionGrantsItemCheckMemberObjectsResponse and sets the default values.
func NewItemPermissionGrantsItemCheckMemberObjectsResponse()(*ItemPermissionGrantsItemCheckMemberObjectsResponse) {
    m := &ItemPermissionGrantsItemCheckMemberObjectsResponse{
        ItemPermissionGrantsItemCheckMemberObjectsPostResponse: *NewItemPermissionGrantsItemCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsItemCheckMemberObjectsResponse(), nil
}
// ItemPermissionGrantsItemCheckMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use checkMemberObjectsPostResponse instead.
type ItemPermissionGrantsItemCheckMemberObjectsResponseable interface {
    ItemPermissionGrantsItemCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
